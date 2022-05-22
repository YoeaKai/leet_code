package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/recover_binary_search_tree"
)

func run() {
	// println(make_array_strictly_increasing.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	d := recover_binary_search_tree.TreeNode{Val: 2}
	c := recover_binary_search_tree.TreeNode{Val: 4, Left: &d}
	b := recover_binary_search_tree.TreeNode{Val: 1}
	a := recover_binary_search_tree.TreeNode{Val: 3, Left: &b, Right: &c}
	recover_binary_search_tree.RecoverTree(&a)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
