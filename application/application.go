package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/spiral_matrix"
)

func run() {
	//println(wildcard_matching.IsMatch("aaa", "*ba"))
	//println(group_anagrams.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	spiral_matrix.SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
