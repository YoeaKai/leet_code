package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/minimum_time_difference"
)

func run() {
	println(minimum_time_difference.FindMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
