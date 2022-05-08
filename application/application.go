package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/binary_tree_inorder_traversal"
)

func run() {
	c := binary_tree_inorder_traversal.TreeNode{Val: 3}
	b := binary_tree_inorder_traversal.TreeNode{Val: 2, Left: &c}
	a := binary_tree_inorder_traversal.TreeNode{Val: 1, Right: &b}
	binary_tree_inorder_traversal.InorderTraversal(&a)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
