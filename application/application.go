package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximum_subarray"
)

func run() {
	//gas_station.CanCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
	println(maximum_subarray.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
