package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/word_search"
)

func run() {
	println(word_search.Exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
