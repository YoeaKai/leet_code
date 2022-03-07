package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/climbing_stairs"
)

func run() {
	//println(unique_paths.UniquePaths(4, 5))
	println(climbing_stairs.ClimbStairs(6))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
