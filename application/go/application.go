package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/toeplitz_matrix"
)

func run() {
	println(toeplitz_matrix.IsToeplitzMatrix(
		[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
	))
	println(toeplitz_matrix.IsToeplitzMatrix(
		[][]int{{18}, {66}},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
