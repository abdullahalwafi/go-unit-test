package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before")
	m.Run()
	fmt.Println("After")
}