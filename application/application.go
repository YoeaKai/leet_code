package main

import (
	"fmt"

	palindrome_partitioning "github.com/YoeaKai/leet_code/topic/palindrome_partitioning"
)

func run() {
	// fmt.Println(permutations.Permute([]int{2, 3, 1}))
	fmt.Println(palindrome_partitioning.Partition("ababbbab"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
