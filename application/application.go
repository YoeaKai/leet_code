package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/permutations_II"
)

func run() {
	//fmt.Println(longest_increasing_path_in_a_matrix.LongestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(permutations_II.PermuteUnique([]int{1, 1, 2, 2}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
