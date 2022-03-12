package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/set_matrix_zeroes"
)

func run() {
	set_matrix_zeroes.SetZeroes([][]int{{0, 1, 2, 6}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
