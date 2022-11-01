package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/where_will_the_ball_fall"
)

func run() {
	println(where_will_the_ball_fall.FindBall(
		[][]int{
			{1, -1, -1, 1, -1, 1, 1, 1, 1},
			{-1, 1, 1, 1, -1, -1, -1, -1, 1},
			{1, -1, -1, -1, -1, 1, -1, 1, 1},
		},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
