package helper

import "testing"

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Imron",
			request: "Imron",
		},
		{
			name:    "Reviady",
			request: "Reviady",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bm.request)
			}
		})
	}
}
