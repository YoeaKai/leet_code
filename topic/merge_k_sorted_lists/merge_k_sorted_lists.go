package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	point := &ListNode{}
	head := point

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			point.Next = list1
			list1 = list1.Next
		} else {
			point.Next = list2
			list2 = list1
			list1 = point.Next.Next
		}
		point = point.Next
	}
	if list1 == nil {
		point.Next = list2
	} else {
		point.Next = list1
	}
	return head.Next
}

func MergeKLists(lists []*ListNode) *ListNode {
	amount := len(lists)

	if amount == 0 {
		return nil
	}

	interval := 1

	for interval < amount {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}
