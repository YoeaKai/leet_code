package remove_duplicates_from_sorted_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	res := prev

	for prev.Next != nil {
		removed := false

		for prev.Next.Next != nil && prev.Next.Next.Val == prev.Next.Val {
			prev.Next = prev.Next.Next
			removed = true
		}

		if removed {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return res.Next
}
