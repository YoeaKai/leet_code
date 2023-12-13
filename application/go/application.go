package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/special_positions_in_a_binary_matrix"
)

func run() {
	println(special_positions_in_a_binary_matrix.NumSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
