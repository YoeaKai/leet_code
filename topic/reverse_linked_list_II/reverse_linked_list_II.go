package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{Val: 0, Next: head}
	head = res

	for i := 0; i < left-1; i++ {
		head = head.Next
	}

	var cur, pre *ListNode = head.Next, nil

	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	head.Next.Next = cur
	head.Next = pre

	return res.Next
}
