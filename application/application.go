package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/best_time_to_buy_and_sell_stock_III"
)

func run() {
	println(best_time_to_buy_and_sell_stock_III.MaxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
