package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/minimum_time_to_make_rope_colorful"
)

func run() {
	println(minimum_time_to_make_rope_colorful.MinCost("bbbaaa", []int{4, 9, 3, 8, 8, 9}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
