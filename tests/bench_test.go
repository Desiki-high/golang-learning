package tests

import (
	"math/rand"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func BenchmarkRandFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Float32()
	}
}

func BenchmarkRandFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Float64()
	}
}
