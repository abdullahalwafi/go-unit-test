package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	// windows
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}

	result := HelloWorld("Abdullah")
	require.Equal(t, "Hello Abdullah", result)
}
