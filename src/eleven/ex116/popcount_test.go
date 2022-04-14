package ex116

import "testing"

const bin = 2 ^ 54

func BenchmarkPopCount(b *testing.B) {

	for i := 0; i < b.N; i++ {
		PopCount(bin)
	}
}
