package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/rotate_list"
)

func run() {

	e := rotate_list.ListNode{Val: 5, Next: nil}
	d := rotate_list.ListNode{Val: 4, Next: &e}
	c := rotate_list.ListNode{Val: 3, Next: &d}
	b := rotate_list.ListNode{Val: 2, Next: &c}
	a := rotate_list.ListNode{Val: 1, Next: &b}
	println(rotate_list.RotateRight(&a, 2))

	// println(length_of_last_word.LengthOfLastWord("   fly me   to   the moon  "))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
