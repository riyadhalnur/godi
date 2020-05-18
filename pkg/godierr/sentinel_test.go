package godierr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentinelErrors(t *testing.T) {
	t.Run("required arguments", func(t *testing.T) {
		err := RequiredArgsError("some error", "some other error")

		assert.Equal(t, 400, err.Code())
		assert.Equal(t, RequiredArgType, err.Type())
		assert.Contains(t, err.Error(), RequiredArgMsg)
	})

	t.Run("invalid arguments", func(t *testing.T) {
		err := InvalidArgsError("some error", "some other error")

		assert.Equal(t, 400, err.Code())
		assert.Equal(t, InvalidArgType, err.Type())
		assert.Contains(t, err.Error(), InvalidArgMsg)
	})
}
