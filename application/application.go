package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/merge_sorted_array"
)

func run() {
	merge_sorted_array.Merge([]int{0}, 0, []int{1}, 1)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
