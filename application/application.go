package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/distinct_subsequences"
)

func run() {
	println(distinct_subsequences.NumDistinct("babgbag", "bag"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
