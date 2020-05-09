package util

import (
	"net/http"
)

// ErrorResponse error structure for errors in http requests
// Code - the error code
// Type - type of error
// Message - full description of error
type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

// Request struct passed in to http handlers
// Contains the full standard http.request
// Path parametrs are mapped to a separate key
// to avoid having to call max.Vars(r)
type Request struct {
	PathParameters map[string]string `json:"pathParameters,omitempty"`

	*http.Request
}

// Response structure of responses
// to return for http handlers
// StatusCode - any valid http status code
// Headers - Customer headers to be attached to the response headers
// Body - the body to be returned as JSON string
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}
