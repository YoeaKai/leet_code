package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/find_k_pairs_with_smallest_sums"
)

func run() {
	//println(unique_paths.UniquePaths(4, 5))
	println(find_k_pairs_with_smallest_sums.KSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
