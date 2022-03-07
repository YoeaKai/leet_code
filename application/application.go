package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/different_ways_to_add_parentheses"
)

func run() {
	//println(unique_paths.UniquePaths(4, 5))
	println(different_ways_to_add_parentheses.DiffWaysToCompute("2*3-4*5"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
