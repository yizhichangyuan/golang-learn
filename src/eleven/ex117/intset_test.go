package main

import (
	"math/rand"
	"testing"
	"time"
)

const n = 1000000
const scale = 100

func gerSlice(rand *rand.Rand) []int {
	var sli []int
	for i := 0; i < n; i++ {
		sli = append(sli, rand.Intn(n*scale))
	}
	return sli
}

func BenchmarkIntSetAdd(b *testing.B) {
	source := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(source))
	sli := gerSlice(rng)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var in IntSet
		for _, n := range sli {
			in.Add(n)
		}
	}
}

func BenchmarkIntSet_Has(b *testing.B) {
	source := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(source))
	sli := gerSlice(rng)

	var in IntSet
	for _, n := range sli {
		in.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, n := range sli {
			if !in.Has(n) {
				b.Errorf("%d not existed, exist acually\n", n)
			}
		}
	}
}
