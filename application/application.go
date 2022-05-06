package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/nonoverlapping_Intervals"
)

func run() {
	println(nonoverlapping_Intervals.EraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
