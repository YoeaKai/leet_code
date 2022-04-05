package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1
func DeleteDuplicates(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		head.Next = DeleteDuplicates(head.Next)

		if head.Val == head.Next.Val {
			return head.Next
		}
	}
	return head
}

// Method 2
func DeleteDuplicates2(head *ListNode) *ListNode {
	store := head
	nowNode := head

	for nowNode != nil {
		if nowNode.Val != store.Val {
			store.Next = nowNode
			store = nowNode
		}
		nowNode = nowNode.Next
	}

	if store != nil {
		store.Next = nil
	}

	return head
}
