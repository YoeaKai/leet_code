package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/my_calendar_I"
)

func run() {
	obj := my_calendar_I.Constructor()
	param_1 := obj.Book(10, 20)
	println(param_1)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
