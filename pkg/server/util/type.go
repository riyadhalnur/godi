package util

type contextKey string

// String will return the string value of the context key
func (c contextKey) String() string {
	return string(c)
}

const (
	// RequestIDKey key for unique request ID attached to
	// all incoming http requests
	RequestIDKey contextKey = "RequestID"
)
