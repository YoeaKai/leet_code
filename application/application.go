package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/boats_to_save_people"
)

func run() {
	println(boats_to_save_people.NumRescueBoats([]int{1, 2}, 3))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
