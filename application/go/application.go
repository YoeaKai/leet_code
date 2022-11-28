package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/find_players_with_zero_or_one_losses"
)

func run() {
	println(find_players_with_zero_or_one_losses.FindWinners(
		[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
