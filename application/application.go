package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/group_anagrams"
)

func run() {
	//gas_station.CanCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
	println(group_anagrams.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
