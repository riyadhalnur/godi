package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDebugMode(t *testing.T) {
	originalFlag := os.Getenv("DEBUG")
	defer func() {
		os.Setenv("DEBUG", originalFlag)
	}()

	os.Setenv("DEBUG", "true")
	assert.Equal(t, true, isDebugMode())
}

func TestStdout(t *testing.T) {
	originalFlag := os.Getenv("DEBUG")
	defer func() {
		os.Setenv("DEBUG", originalFlag)
	}()

	os.Setenv("DEBUG", "true")

	logger.Info("info will be logged")
	logger.Debug("debug will be logged")

	// Output:
	// {"level":"INFO","timestamp":"2020-05-13T00:06:47.460+0800","caller":"logger/logger_test.go:35","message":"info will be logged"}
	// {"level":"DEBUG","timestamp":"2020-05-13T00:06:47.460+0800","caller":"logger/logger_test.go:36","message":"debug will be logged"}

	os.Setenv("DEBUG", "false")

	logger.Info("info will be logged")
	logger.Debug("debug will not be logged")

	// Output:
	// {"level":"INFO","timestamp":"2020-05-13T00:06:47.460+0800","caller":"logger/logger_test.go:35","message":"info will be logged"}
}
