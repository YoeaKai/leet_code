package main

import (
	"testing"
	"topic/benchmark"
)

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark.SolveNQueens(10)
	}
}
