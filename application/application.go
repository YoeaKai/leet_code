package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/text_justification"
)

func run() {
	println(text_justification.FullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
