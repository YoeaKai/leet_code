package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/make_array_strictly_increasing"
)

func run() {
	println(make_array_strictly_increasing.MakeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
