package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/godi/pkg/logger"

	"github.com/godi/pkg/middleware"
	"github.com/godi/pkg/server/util"

	"github.com/gorilla/mux"
)

const (
	staticPathPrefix string = "/static"
)

var (
	listenPort = "3000"
)

// Server holds the configurations,
// routes, and the middlewares
// to mount when an instance of Server starts
type Server struct {
	config      *Config
	routers     []util.Route
	middlewares []mux.MiddlewareFunc
}

// NewServer returns a new instance of Server
// with passed in configation
func NewServer(cfg *Config) *Server {
	return &Server{
		config: cfg,
	}
}

// Listen will handle incoming HTTP requests
// Blocks until an interrupt is received
func (s *Server) Listen() error {
	if s.config.Timeout == 0 {
		return errors.New("timeout configuration is a required value")
	}

	if s.config.Port != "" {
		listenPort = s.config.Port
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", listenPort),
		WriteTimeout: time.Duration(s.config.Timeout) * time.Second,
		ReadTimeout:  time.Duration(s.config.Timeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.Timeout) * time.Second,
		Handler:      s.mountRoutes(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalf("Unable to start server err=%v", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// wait for active connections to finish their jobs
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Timeout)*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

// Close shutdowns the server immediately
func (s *Server) Close() error {
	return s.Close()
}

// AddRoutes appends the list of routes to mount
func (s *Server) AddRoutes(routes ...util.Route) {
	s.routers = append(s.routers, routes...)
}

// AddMiddlewares appends the middleware(s) to mount
// Middlewares should be ordered according to their functionality
func (s *Server) AddMiddlewares(middleware ...mux.MiddlewareFunc) {
	s.middlewares = append(s.middlewares, middleware...)
}

func (s *Server) handleHTTP(handler util.APIHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// reuse context attached to request and pass in to handler
		ctx := r.Context()
		params := mux.Vars(r)

		req := &util.Request{
			PathParameters: params,
			Request:        r,
		}

		res, err := handler(ctx, req)
		if err != nil {
			util.ErrorJSON(w, &util.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		util.RespondJSON(w, res)
	}
}

func (s *Server) mountRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	if s.config.StaticDir != "" {
		staticAbsPath, err := filepath.Abs(s.config.StaticDir)
		if err != nil {
			logger.Errorf("Unable to mount static directory err=%v", err.Error())
		}

		// serve static files at /static path from /static folder unless specified otherwise
		fs := http.FileServer(http.Dir(staticAbsPath))
		router.PathPrefix(staticPathPrefix).Handler(http.StripPrefix(staticPathPrefix, fs))
	}

	// attach default middlewares
	// order matters for middleware
	router.Use(middleware.RequestID)

	// mount the health enpoint. useful for Kubernetes integration
	router.Name("health").Path("/health").HandlerFunc(s.handleHTTP(healthCheckHandler)).Methods(http.MethodGet)

	// mount middlewares
	for _, mw := range s.middlewares {
		router.Use(mw)
	}

	// mount routes
	for _, route := range s.routers {
		router.Name(route.Name).Path(route.Path).HandlerFunc(s.handleHTTP(route.Handler)).Methods(route.Method)
	}

	return router
}

func healthCheckHandler(ctx context.Context, req *util.Request) (*util.Response, error) {
	return &util.Response{
		StatusCode: http.StatusOK,
		Body:       "ok",
	}, nil
}
