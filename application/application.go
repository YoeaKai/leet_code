package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/reverse_linked_list_II"
)

func run() {
	// b := reverse_linked_list_II.ListNode{Val: 5, Next: nil}
	// a := reverse_linked_list_II.ListNode{Val: 3, Next: &b}
	// reverse_linked_list_II.ReverseBetween(&a, 1, 2)
	e := reverse_linked_list_II.ListNode{Val: 5, Next: nil}
	d := reverse_linked_list_II.ListNode{Val: 4, Next: &e}
	c := reverse_linked_list_II.ListNode{Val: 3, Next: &d}
	b := reverse_linked_list_II.ListNode{Val: 2, Next: &c}
	a := reverse_linked_list_II.ListNode{Val: 1, Next: &b}
	reverse_linked_list_II.ReverseBetween(&a, 1, 5)
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
