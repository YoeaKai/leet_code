package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/convert_BST_to_greater_tree"
)

func run() {
	e := convert_BST_to_greater_tree.TreeNode{Val: 5, Right: nil, Left: nil}
	d := convert_BST_to_greater_tree.TreeNode{Val: 2, Right: nil, Left: nil}
	c := convert_BST_to_greater_tree.TreeNode{Val: 6, Right: nil, Left: &e}
	b := convert_BST_to_greater_tree.TreeNode{Val: 1, Right: &d, Left: nil}
	a := convert_BST_to_greater_tree.TreeNode{Val: 4, Right: &c, Left: &b}
	println(convert_BST_to_greater_tree.ConvertBST(&a))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
