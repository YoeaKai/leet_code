package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/contiguous_array"
)

func run() {
	println(contiguous_array.FindMaxLength([]int{0, 1, 1, 0, 1, 1, 1, 0}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
