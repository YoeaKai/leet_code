package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/coin_change_2"
)

func run() {
	println(coin_change_2.Change(5, []int{1, 2, 5}))
	// println(length_of_last_word.LengthOfLastWord("   fly me   to   the moon  "))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
