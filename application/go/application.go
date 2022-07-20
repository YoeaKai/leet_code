package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/number_of_matching_subsequences"
)

func run() {
	println(number_of_matching_subsequences.NumMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
