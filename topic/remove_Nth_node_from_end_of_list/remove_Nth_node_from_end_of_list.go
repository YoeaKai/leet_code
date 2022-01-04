package remove_Nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := &ListNode{Val: -1, Next: head}

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Val == -1 {
		return head.Next
	}

	slow.Next = slow.Next.Next
	return head
}
