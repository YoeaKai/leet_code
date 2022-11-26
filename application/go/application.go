package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/maximum_profit_in_job_scheduling"
)

func run() {
	println(maximum_profit_in_job_scheduling.JobScheduling(
		[]int{1, 2, 3, 3},
		[]int{3, 4, 5, 6},
		[]int{50, 10, 40, 70},
	))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
