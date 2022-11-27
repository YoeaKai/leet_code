package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/arithmetic_slices_II_subsequence"
)

func run() {
	println(arithmetic_slices_II_subsequence.NumberOfArithmeticSlices(
		[]int{2, 2, 3, 4},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
