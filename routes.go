package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func routes() []*Route {
	return []*Route{
		&Route{
			"getUser",
			http.MethodGet,
			"/user/{id}",
			GetUser,
		},
		&Route{
			"createUser",
			http.MethodPost,
			"/user",
			CreateUser,
		},
	}
}
