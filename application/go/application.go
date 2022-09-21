package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/sum_of_even_numbers_after_queries"
)

func run() {
	println(sum_of_even_numbers_after_queries.SumEvenAfterQueries([]int{-1, 3, -3, 9, -6, 8, -5},
		[][]int{{-5, 1}, {10, 2}, {-6, 3}, {3, 2}, {9, 5}, {7, 5}, {8, 0}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
