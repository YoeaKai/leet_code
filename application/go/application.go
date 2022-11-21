package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/nearest_exit_from_entrance_in_maze"
)

func run() {
	println(nearest_exit_from_entrance_in_maze.NearestExit(
		[][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
		[]int{1, 0},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
