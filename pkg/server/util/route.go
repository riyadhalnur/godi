package util

import (
	"context"
)

// APIHandlerFunc is the signature API controllers must implement
// to be used as part of the server instance
type APIHandlerFunc func(ctx context.Context, req *Request) (res *Response, err error)

// Route defines the properties of an API route
// to mount on the server
type Route struct {
	Name    string
	Path    string
	Method  string
	Handler APIHandlerFunc
}
