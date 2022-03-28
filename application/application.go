package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/find_bottom_left_tree_value"
)

func run() {
	a := find_bottom_left_tree_value.TreeNode{Val: 1}
	b := find_bottom_left_tree_value.TreeNode{Val: 3}
	c := find_bottom_left_tree_value.TreeNode{Val: 2, Left: &a, Right: &b}
	println(find_bottom_left_tree_value.FindBottomLeftValue(&c))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
