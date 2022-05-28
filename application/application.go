package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/target_sum"
)

func run() {
	println(target_sum.FindTargetSumWays([]int{5, 4, 3, 2, 1}, 5))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
