package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximum_product_of_word_lengths"
)

func run() {
	println(maximum_product_of_word_lengths.MaxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
