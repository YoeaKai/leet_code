package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/image_smoother"
)

func run() {
	println(image_smoother.ImageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
