package main

import (
	"testing"

	"github.com/YoeaKai/leet_code/topic/n_queens"
)

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n_queens.SolveNQueens(10)
	}
}
