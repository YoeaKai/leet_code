package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximum_depth_of_binary_tree"
)

func run() {
	e := maximum_depth_of_binary_tree.TreeNode{Val: 5}
	d := maximum_depth_of_binary_tree.TreeNode{Val: 2}
	c := maximum_depth_of_binary_tree.TreeNode{Val: 4, Left: &d}
	b := maximum_depth_of_binary_tree.TreeNode{Val: 1, Right: &e}
	a := maximum_depth_of_binary_tree.TreeNode{Val: 3, Left: &b, Right: &c}
	println(maximum_depth_of_binary_tree.MaxDepth(&a))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
