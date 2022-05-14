package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/interleaving_string"
)

func run() {
	// println(interleaving_string.IsInterleave("a", "", "c"))
	println(interleaving_string.IsInterleave("ab", "bc", "babc"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
