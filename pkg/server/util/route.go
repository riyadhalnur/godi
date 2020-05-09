package util

import (
	"context"
)

type APIHandlerFunc func(ctx context.Context, req *Request) (res *Response, err error)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler APIHandlerFunc
}
