package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/edit_distance"
)

func run() {
	println(edit_distance.MinDistance("horse", "ros"))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
