package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/jump_game"
)

func run() {
	//println(wildcard_matching.IsMatch("aaa", "*ba"))
	println(jump_game.CanJump([]int{1, 2}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
