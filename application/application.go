package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/two_city_scheduling"
)

func run() {
	println(two_city_scheduling.TwoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
