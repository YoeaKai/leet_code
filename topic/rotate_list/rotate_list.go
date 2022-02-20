package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	tail := head
	len := 1

	for tail.Next != nil {
		tail = tail.Next
		len++
	}

	tail.Next = head

	k %= len

	if k != 0 {
		for i := 0; i < len-k; i++ {
			tail = tail.Next
		}
	}

	head = tail.Next
	tail.Next = nil

	return head
}
