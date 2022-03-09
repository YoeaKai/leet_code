package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/split_array_largest_sum"
)

func run() {
	println(split_array_largest_sum.SplitArray([]int{7, 2, 5, 10, 8}, 2))
	//println(split_array_largest_sum.SplitArray([]int{7, 2, 5, 10, 8}, 4))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
