package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/most_stones_removed_with_same_row_or_column"
)

func run() {
	println(most_stones_removed_with_same_row_or_column.RemoveStones(
		[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
