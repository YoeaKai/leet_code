package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/destination_city"
)

func run() {
	println(destination_city.DestCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
