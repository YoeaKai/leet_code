package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/longest_palindrome_by_concatenating_two_letter_words"
)

func run() {
	println(longest_palindrome_by_concatenating_two_letter_words.LongestPalindrome2(
		[]string{"bb", "bb"},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
