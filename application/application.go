package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximize_distance_to_closest_person"
)

func run() {
	//rotate_image.Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	println(maximize_distance_to_closest_person.MaxDistToClosest([]int{0, 0, 1, 0, 0, 1}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
