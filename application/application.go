package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/unique_paths_II"
)

func run() {
	// unique_paths_II.UniquePathsWithObstacles([][]int{{0, 1}, {0, 0}})
	unique_paths_II.UniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
