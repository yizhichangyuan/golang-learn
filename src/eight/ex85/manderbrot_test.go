package mandelbrot

import "testing"

func benchmarkCurrentRender(b *testing.B, workers int) {
	for i := 0; i < b.N; i++ {
		Render(workers)
	}
}

func Benchmark1(b *testing.B) {
	benchmarkCurrentRender(b, 8)
}
