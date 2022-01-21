package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/gas_station"
)

func run() {
	//gas_station.CanCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
	gas_station.CanCompleteCircuit([]int{1, 2, 3, 4, 3, 2, 4, 1, 5, 3, 2, 4}, []int{1, 1, 1, 3, 2, 4, 3, 6, 7, 4, 3, 1})
	//println(maximize_distance_to_closest_person.MaxDistToClosest([]int{0, 0, 1, 0, 0, 1}))
	//construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
