package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/same_tree"
)

func run() {
	b := same_tree.TreeNode{Val: 1}
	a := same_tree.TreeNode{Val: 1}
	println(same_tree.IsSameTree(&a, &b))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
