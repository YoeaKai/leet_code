package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/earliest_possible_day_of_full_bloom"
)

func run() {
	println(earliest_possible_day_of_full_bloom.EarliestFullBloom(
		[]int{1, 4, 3},
		[]int{2, 3, 1},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
