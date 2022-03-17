package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/sum_of_distances_in_tree"
)

func run() {
	println(sum_of_distances_in_tree.SumOfDistancesInTree(5, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
