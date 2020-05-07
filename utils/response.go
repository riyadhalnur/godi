package utils

import (
	"net/http"
)

// ErrorResponse defines the error structure
// for http responses
// Code - the error code
// Type - type of error
// Message - ful description of error
type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

// Response defines the structure of responses
// for all http reponses.
// StatusCode - any valid http status code
// Headers - Customer headers to be attached to the response headers
// Body - the body to be returned as byte
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       []byte            `json:"body,omitempty"`
}

// RespondJSON returns fully formed JSON responses
// for http requests
func RespondJSON(w http.ResponseWriter, response *Response) {
	w.Header().Set("Content-Type", "application/json")

	if len(response.Headers) != 0 {
		for k, v := range response.Headers {
			// set any custom headers passed in
			w.Header().Set(k, v)
		}
	}

	w.WriteHeader(response.StatusCode)
	w.Write(response.Body)
}
