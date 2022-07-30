package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/word_subsets"
)

func run() {
	println(word_subsets.WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
