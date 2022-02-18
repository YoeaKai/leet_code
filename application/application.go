package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/insert_interval"
)

func run() {
	println(insert_interval.Insert([][]int{{1, 5}}, []int{0, 0}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
