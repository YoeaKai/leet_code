package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/minimum_swaps_to_make_sequences_increasing"
)

func run() {
	println(minimum_swaps_to_make_sequences_increasing.MinSwap([]int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}, []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
