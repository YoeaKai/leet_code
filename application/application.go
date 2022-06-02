package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/partition_equal_subset_sum"
)

func run() {
	// println(partition_equal_subset_sum.CanPartition([]int{3, 3, 3, 4, 5}))
	println(partition_equal_subset_sum.CanPartition([]int{2, 2, 3, 5}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
