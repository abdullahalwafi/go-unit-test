package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Abdullah")
	assert.Equal(t, "Hello Abdullah", result)
	fmt.Println("dieksekusi")
}
