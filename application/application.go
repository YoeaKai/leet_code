package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/longest_increasing_path_in_a_matrix"
)

func run() {
	longest_increasing_path_in_a_matrix.LongestIncreasingPathII([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})
	//longest_increasing_path_in_a_matrix.LongestIncreasingPathII([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})
	//longest_increasing_path_in_a_matrix.LongestIncreasingPathII([][]int{{6, 8}, {7, 2}})
	//println(maximize_distance_to_closest_person.MaxDistToClosest([]int{0, 0, 1, 0, 0, 1}))
	//construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
