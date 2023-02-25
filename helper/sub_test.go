package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Abdullah", func(t *testing.T) {
		result := HelloWorld("Abdullah")
		require.Equal(t, "Hello Abdullah", result)
	})
	t.Run("Al Wafi", func(t *testing.T) {
		result := HelloWorld("Al Wafi")
		require.Equal(t, "Hello Al Wafi", result)
	})
}
