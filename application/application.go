package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/flatten_a_multilevel_doubly_linked_list"
)

func run() {
	c := flatten_a_multilevel_doubly_linked_list.Node{Val: 3, Next: nil, Child: nil}
	b := flatten_a_multilevel_doubly_linked_list.Node{Val: 2, Next: nil, Child: &c}
	a := flatten_a_multilevel_doubly_linked_list.Node{Val: 1, Next: nil, Child: &b}
	// e := flatten_a_multilevel_doubly_linked_list.Node{Val: 5, Next: nil, Child: nil}
	// d := flatten_a_multilevel_doubly_linked_list.Node{Val: 4, Next: &e, Child: nil}
	// c := flatten_a_multilevel_doubly_linked_list.Node{Val: 3, Next: nil, Child: nil}
	// b := flatten_a_multilevel_doubly_linked_list.Node{Val: 2, Next: &c, Child: &d}
	// a := flatten_a_multilevel_doubly_linked_list.Node{Val: 1, Next: &b, Child: nil}
	println(flatten_a_multilevel_doubly_linked_list.Flatten(&a))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
