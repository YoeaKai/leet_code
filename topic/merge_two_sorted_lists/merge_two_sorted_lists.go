package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	point := &ListNode{}
	head := point

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			point.Next = l1
			l1 = l1.Next
		} else {
			point.Next = l2
			l2 = l1
			l1 = point.Next.Next
		}
		point = point.Next
	}
	if l1 == nil {
		point.Next = l2
	} else {
		point.Next = l1
	}
	return head.Next
}

// Method 2
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	head := l1
	storeNode := l2
	isL1 := true

	if l1.Val > l2.Val {
		head = l2
		storeNode = l1
		isL1 = false
	}

	for l1 != nil && l2 != nil {
		if isL1 {
			if l1.Next == nil {
				l1.Next = storeNode
				break
			}
			if l1.Next.Val < storeNode.Val {
				l1 = l1.Next
			} else {
				tmpNode := storeNode
				storeNode = l1.Next
				l1.Next = tmpNode
				l1 = storeNode
				isL1 = false
			}
		} else {
			if l2.Next == nil {
				l2.Next = storeNode
				break
			}
			if l2.Next.Val < storeNode.Val {
				l2 = l2.Next
			} else {
				tmpNode := storeNode
				storeNode = l2.Next
				l2.Next = tmpNode
				l2 = storeNode
				isL1 = true
			}
		}
	}

	return head
}
