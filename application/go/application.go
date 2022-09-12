package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/bag_of_tokens"
)

func run() {
	println(bag_of_tokens.BagOfTokensScore([]int{71, 55, 82}, 54))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
