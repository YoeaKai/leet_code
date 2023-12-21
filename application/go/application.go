package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/widest_vertical_area_between_two_points_containing_no_points"
)

func run() {
	println(widest_vertical_area_between_two_points_containing_no_points.MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
