package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/product_of_array_except_self"
)

func run() {
	//rotate_image.Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	println(product_of_array_except_self.ProductExceptSelf([]int{0, 0, 0}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
