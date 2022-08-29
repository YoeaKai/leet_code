package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/number_of_islands"
)

func run() {
	println(number_of_islands.NumIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
