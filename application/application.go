package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/search_a_2D_matrix"
)

func run() {
	println(search_a_2D_matrix.SearchMatrix([][]int{{1, 1}}, 2))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
