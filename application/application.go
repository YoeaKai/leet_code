package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/shortest_unsorted_continuous_subarray"
)

func run() {
	println(shortest_unsorted_continuous_subarray.FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	// println(length_of_last_word.LengthOfLastWord("   fly me   to   the moon  "))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
