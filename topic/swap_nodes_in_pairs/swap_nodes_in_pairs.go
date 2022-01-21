package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = SwapPairs(node.Next)
	node.Next = head
	return node
}

// Method 2
func SwapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next.Next
	result := head.Next
	result.Next = head
	head.Next = tmp
	pos := tmp

	for pos != nil && pos.Next != nil {
		tmp = pos.Next.Next
		head.Next = pos.Next
		head.Next.Next = pos
		pos.Next = tmp
		head = pos
		pos = tmp
	}

	return result
}
