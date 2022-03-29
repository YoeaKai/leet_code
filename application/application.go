package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/longest_common_subsequence"
)

func run() {
	println(longest_common_subsequence.LongestCommonSubsequence("bcbbd", "abcbad"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
