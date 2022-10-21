package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/contains_duplicate_II"
)

func run() {
	println(contains_duplicate_II.ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
