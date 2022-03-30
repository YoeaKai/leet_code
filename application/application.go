package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/search_in_rotated_sorted_array_II"
)

func run() {
	println(search_in_rotated_sorted_array_II.Search([]int{2, 5, 6, 0, 0, 1, 2}, 1))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
