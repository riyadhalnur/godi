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
	// {"level":"INFO","timestamp":"2020-10-27T17:52:44.121+0800","caller":"testing/testing.go:1123","message":"info will be logged"}
    //{"level":"DEBUG","timestamp":"2020-10-27T17:52:44.121+0800","caller":"testing/testing.go:1123","message":"debug will be logged"}

	os.Setenv("DEBUG", "false")

	logger.Info("info will be logged")
	logger.Debug("debug will not be logged")

	// Output:
	// {"level":"INFO","timestamp":"2020-10-27T17:52:44.121+0800","caller":"testing/testing.go:1123","message":"info will be logged"}
}
