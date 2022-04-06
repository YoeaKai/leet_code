package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/remove_duplicates_from_sorted_list_II"
)

func run() {
	e := remove_duplicates_from_sorted_list_II.ListNode{Val: 3, Next: nil}
	d := remove_duplicates_from_sorted_list_II.ListNode{Val: 3, Next: &e}
	c := remove_duplicates_from_sorted_list_II.ListNode{Val: 2, Next: &d}
	b := remove_duplicates_from_sorted_list_II.ListNode{Val: 1, Next: &c}
	a := remove_duplicates_from_sorted_list_II.ListNode{Val: 1, Next: &b}
	println(remove_duplicates_from_sorted_list_II.DeleteDuplicates(&a))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
