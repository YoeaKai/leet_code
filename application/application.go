package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/symmetric_tree"
)

func run() {
	// println(make_array_strictly_increasing.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	println(symmetric_tree.IsSymmetric(nil))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
