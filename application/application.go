package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/minimum_number_of_steps_to_make_two_strings_anagram"
)

func run() {
	println(minimum_number_of_steps_to_make_two_strings_anagram.MinSteps("bab", "aba"))
	// println(jump_game.CanJump([]int{1, 2}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
