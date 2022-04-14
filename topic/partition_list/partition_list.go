package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	var lessStart, moreStart ListNode
	less, more := &lessStart, &moreStart

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			more.Next = head
			more = more.Next
		}
		head = head.Next
	}

	more.Next = nil
	less.Next = moreStart.Next

	return lessStart.Next
}
