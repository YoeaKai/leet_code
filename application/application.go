package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/triangle"
)

func run() {
	println(triangle.MinimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
