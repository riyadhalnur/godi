package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextKey(t *testing.T) {
	const randomKey contextKey = "hello"
	assert.Equal(t, "hello", randomKey.String())
}
