package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/rotate_image"
)

func run() {
	//fmt.Println(longest_increasing_path_in_a_matrix.LongestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	rotate_image.Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
