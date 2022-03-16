package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/increasing_triplet_subsequence"
)

func run() {
	println(increasing_triplet_subsequence.IncreasingTriplet([]int{1, 1, -2, 6}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
