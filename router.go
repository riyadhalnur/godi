package main

import (
	"net/http"

	"github.com/simple-go-server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// NewRouter returns a mux router initialised with
// basic middleware and static and api routes mounted
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// serve static files at /static path from /static folder unless specified otherwise
	fs := http.FileServer(http.Dir("." + STATIC))
	router.PathPrefix(STATIC).Handler(http.StripPrefix(STATIC, fs))

	api := router.PathPrefix("/api").Subrouter()
	// register middlewares
	api.Use(middleware.Logging)
	api.Use(middleware.Authenticate)
	api.Use(handlers.CompressHandler)
	api.Use(handlers.CORS())

	// mount any routes registered in route list
	for _, route := range routes() {
		api.HandleFunc(route.Path, route.Handler).Name(route.Name).Methods(route.Method)
	}

	return router
}
