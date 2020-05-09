package util

import (
	"encoding/json"
	"net/http"
)

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
	w.Write([]byte(response.Body))
}

func ErrorJSON(w http.ResponseWriter, err *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
