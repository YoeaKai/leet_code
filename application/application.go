package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows"
)

func run() {
	println(find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows.KthSmallest([][]int{{1, 10, 10}, {1, 4, 5}}, 7))
	// println(find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows.KthSmallest([][]int{{2, 5, 6, 11, 11, 14, 14}, {2, 3, 6}}, 7))
	// println(find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows.KthSmallest([][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
