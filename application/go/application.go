package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/pascals_triangle"
)

func run() {
	println(pascals_triangle.Generate(10))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
