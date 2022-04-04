package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/brick_wall"
)

func run() {
	println(brick_wall.LeastBricks([][]int{{100000000}, {100000000}, {100000000}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
