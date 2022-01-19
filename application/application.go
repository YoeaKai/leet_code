package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/construct_binary_tree_from_preorder_and_inorder_traversal"
)

func run() {
	//rotate_image.Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	//println(maximize_distance_to_closest_person.MaxDistToClosest([]int{0, 0, 1, 0, 0, 1}))
	construct_binary_tree_from_preorder_and_inorder_traversal.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
