package godierr

import "fmt"

// Error wraps a Go error
// with more human friendly information.
// Based on the fact that runtime errors should stay
// within the system - logs etc and vanilla messages
// be propragated client side
type Error struct {
	error

	code    int
	t       string
	message string
}

// Error returns a formatted version
// of the error message and the original
// error passed in
func (e *Error) Error() string {
	msg := fmt.Sprintf("%s", e.message)
	if e.error != nil {
		msg = fmt.Sprintf("%s due to %s", msg, e.error.Error())
	}
	return msg
}

// Code returns the error code
func (e *Error) Code() int {
	return e.code
}

// Type returns the error type
func (e *Error) Type() string {
	return e.t
}

// Message returns the error message
func (e *Error) Message() string {
	return e.message
}

// New returns an Error object. Takes the error code,
// type, message, and the original error, if any
func New(code int, t, msg string, err error) *Error {
	return &Error{
		code:    code,
		t:       t,
		message: msg,
		error:   err,
	}
}
