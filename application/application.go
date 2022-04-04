package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/kth_smallest_number_in_multiplication_table"
)

func run() {
	println(kth_smallest_number_in_multiplication_table.FindKthNumber(3, 3, 5))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
