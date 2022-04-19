package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/zero_one_matrix"
)

func run() {
	println(zero_one_matrix.UpdateMatrix([][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
