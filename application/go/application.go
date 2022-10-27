package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/image_overlap"
)

func run() {
	println(image_overlap.LargestOverlap(
		[][]int{
			{1, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 1},
		}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
