package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/course_schedule"
)

func run() {
	println(course_schedule.CanFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
