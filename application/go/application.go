package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/interview_MT_2022_10"
)

func run() {
	println(interview_MT_2022_10.Q3([]string{
		".##.#",
		"#.#..",
		"#...#",
		"#.##.",
	}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
