package helper

import "testing"

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Imron", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Imron")
		}
	})
	b.Run("Reviady", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Reviady")
		}
	})
}
