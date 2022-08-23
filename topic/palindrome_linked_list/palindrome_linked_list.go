package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	mid, tail := head, head
	var pre, tmp *ListNode

	for tail != nil && tail.Next != nil {
		tail = tail.Next.Next

		// Reverse half of list.
		pre = mid
		mid = mid.Next
		pre.Next = tmp
		tmp = pre
	}

	left, right := pre, mid

	if tail != nil {
		// Add a number of nodes if the length of the list is even.
		right = right.Next
	}

	for left != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}

	return true
}
