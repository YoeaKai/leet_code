package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/deepest_leaves_sum"
)

func run() {
	// println(make_array_strictly_increasing.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	e := deepest_leaves_sum.TreeNode{Val: 5}
	d := deepest_leaves_sum.TreeNode{Val: 2}
	c := deepest_leaves_sum.TreeNode{Val: 4, Left: &d}
	b := deepest_leaves_sum.TreeNode{Val: 1, Right: &e}
	a := deepest_leaves_sum.TreeNode{Val: 3, Left: &b, Right: &c}
	deepest_leaves_sum.DeepestLeavesSum(&a)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
