package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/continuous_subarray_sum"
)

func run() {
	println(continuous_subarray_sum.CheckSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
