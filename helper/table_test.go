package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Imron",
			input:    "Imron",
			expected: "Hello Imron",
		},
		{
			name:     "Reviady",
			input:    "Reviady",
			expected: "Hello Reviady",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}