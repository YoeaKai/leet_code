package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/word_search_II"
)

func run() {
	println(word_search_II.FindWords(
		[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
