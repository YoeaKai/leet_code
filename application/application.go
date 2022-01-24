package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/wildcard_matching"
)

func run() {
	println(wildcard_matching.IsMatch("aaa", "*ba"))
	//println(group_anagrams.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
