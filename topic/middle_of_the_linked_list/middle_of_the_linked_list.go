package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	mid := head
	tail := head

	for tail != nil && tail.Next != nil {
		mid = mid.Next
		tail = tail.Next.Next
	}

	return mid
}
