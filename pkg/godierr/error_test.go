package godierr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGodiError(t *testing.T) {
	msg := "this is an error"
	formattedErr := "this is an error due to some error"

	origErr := errors.New("some error")
	err := New(600, "RANDOM_ERROR", msg, origErr)

	assert.Equal(t, msg, err.Message())
	assert.Equal(t, formattedErr, err.Error())
}
