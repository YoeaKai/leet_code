package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximum_product_of_two_elements_in_an_array"
)

func run() {
	println(maximum_product_of_two_elements_in_an_array.MaxProduct([]int{2, 5, 3, 9, 5, 3}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
