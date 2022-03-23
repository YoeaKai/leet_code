package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/remove_duplicates_from_sorted_array_II"
)

func run() {
	println(remove_duplicates_from_sorted_array_II.RemoveDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 2, 2, 3, 4}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
