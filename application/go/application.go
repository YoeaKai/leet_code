package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/pacific_atlantic_water_flow"
)

func run() {
	println(pacific_atlantic_water_flow.PacificAtlantic(
		[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
