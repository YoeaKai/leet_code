package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/minimum_genetic_mutation"
)

func run() {
	println(minimum_genetic_mutation.MinMutation(
		"AACCGGTT",
		"AACCGGTA",
		[]string{"AACCGGTA"},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
