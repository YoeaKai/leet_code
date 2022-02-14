package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/largest_rectangle_in_histogram"
)

func run() {
	println(largest_rectangle_in_histogram.LargestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0}))
	// println(jump_game.CanJump([]int{1, 2}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
