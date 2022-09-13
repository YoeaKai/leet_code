package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/utf_8_validation"
)

func run() {
	println(utf_8_validation.ValidUtf8([]int{197, 130, 1}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
