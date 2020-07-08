package godierr

import (
	"fmt"
	"strings"
)

const (
	// RequiredArgType is the constant error "type" for required arguments
	RequiredArgType string = "REQUIRED_ARGUMENT"
	// InvalidArgType is the constant error "type" for invalid arguments
	InvalidArgType string = "INVALID_ARGUMENT"

	// RequiredArgMsg is the constant extended error "message" for required arguments
	RequiredArgMsg string = "missing required argument(s)"
	// InvalidArgMsg is the constant extended error "message" for invalid arguments
	InvalidArgMsg string = "invalid argument(s) passed in"
)

// RequiredArgsError forms standardised required arguments
// error type. Takes a list of arguments
func RequiredArgsError(args ...string) *Error {
	msg := fmt.Sprintf("%s: %s", RequiredArgMsg, strings.Join(args, ", "))
	return New(400, RequiredArgType, msg, nil)
}

// InvalidArgsError forms standardised invalid arguments
// error type. Takes a list of arguments
func InvalidArgsError(args ...string) *Error {
	msg := fmt.Sprintf("%s: %s", InvalidArgMsg, strings.Join(args, ", "))
	return New(400, InvalidArgType, msg, nil)
}
