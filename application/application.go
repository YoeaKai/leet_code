package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximal_rectangle"
)

func run() {
	println(maximal_rectangle.MaximalRectangle([][]byte{{'0', '1', '1', '0'}, {'0', '0', '1', '0'}, {'0', '0', '1', '0'}, {'0', '0', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}, {'1', '1', '1', '1'}}))
	// println(jump_game.CanJump([]int{1, 2}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
