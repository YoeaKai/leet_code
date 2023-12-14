package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/difference_between_ones_and_zeros_in_row_and_column"
)

func run() {
	println(difference_between_ones_and_zeros_in_row_and_column.OnesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
