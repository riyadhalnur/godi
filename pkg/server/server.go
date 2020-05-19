package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/godi/pkg/godierr"

	"github.com/godi/pkg/logger"

	"github.com/godi/pkg/middleware"
	"github.com/godi/pkg/server/util"

	"github.com/gorilla/mux"
)

const (
	staticPathPrefix string = "/static"
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
		return godierr.RequiredArgsError("timeout")
	}

	if s.config.Port == "" {
		return godierr.RequiredArgsError("port")
	}

	listenPort := s.config.Port
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", listenPort),
		WriteTimeout: time.Duration(s.config.Timeout) * time.Second,
		ReadTimeout:  time.Duration(s.config.Timeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.Timeout) * time.Second,
		Handler:      s.mountRoutes(),
	}

	go func() {
		logger.Infof("Server listening on port %s", listenPort)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalf("Unable to start server err=%v", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Debugf("Interrupt received. Starting shutdown")

	// wait for active connections to finish their jobs
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Timeout)*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		logger.Debugf("Could not shutdown server gracefully err=%v", err)
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
		start := time.Now()
		// reuse context attached to request and pass in to handler
		ctx := r.Context()
		params := mux.Vars(r)

		req := &util.Request{
			PathParameters: params,
			Request:        r,
		}

		logger.Info("Incoming HTTP request",
			"method",
			req.Method,
			"path",
			req.URL.Path,
			"query",
			req.URL.RawQuery,
			"requestId",
			ctx.Value(util.RequestIDKey).(string),
			"ip",
			req.RemoteAddr,
		)

		res, err := handler(ctx, req)
		if err != nil {
			if godiErr, ok := err.(*godierr.Error); ok {
				logger.Error("HTTP handler returned an error",
					"code",
					godiErr.Code(),
					"type",
					godiErr.Type(),
					"error",
					godiErr.Error(),
					"requestId",
					ctx.Value(util.RequestIDKey).(string),
					"latency",
					time.Since(start).String(),
				)

				util.ErrorJSON(w, &util.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Type:    godiErr.Type(),
					Message: godiErr.Message(),
				})
				return
			}

			logger.Error("HTTP handler returned an error",
				"error",
				err.Error(),
				"requestId",
				ctx.Value(util.RequestIDKey).(string),
				"latency",
				time.Since(start).String(),
			)

			util.ErrorJSON(w, &util.ErrorResponse{
				Code: http.StatusInternalServerError,
			})
			return
		}

		logger.Info("Handled HTTP request",
			"method",
			req.Method,
			"path",
			req.URL.Path,
			"status",
			res.StatusCode,
			"requestId",
			ctx.Value(util.RequestIDKey).(string),
			"latency",
			time.Since(start).String(),
		)

		util.RespondJSON(w, res)
	}
}

func (s *Server) mountRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	if s.config.StaticDir != "" {
		staticAbsPath, err := filepath.Abs(s.config.StaticDir)
		if err != nil {
			logger.Errorf("Unable to read absolute path to static directory err=%v", err.Error())
		}

		// serve static files at /static path
		fs := http.FileServer(http.Dir(staticAbsPath))
		router.PathPrefix(staticPathPrefix).Handler(http.StripPrefix(staticPathPrefix, fs))
	}

	router.Use(middleware.RequestID)

	// mount the health enpoint. useful for Kubernetes integration
	router.Name("health").Path("/health").HandlerFunc(s.handleHTTP(healthCheckHandler)).Methods(http.MethodGet)

	logger.Debug("Mounting middlewares")
	for _, mw := range s.middlewares {
		router.Use(mw)
	}

	for _, route := range s.routers {
		logger.Debug("Mounting route", "name", route.Name, "path", route.Path, "method", route.Method)
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
