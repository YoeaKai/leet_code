package reverse_nodes_in_kGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	nowNode := head
	tmpNode := head.Next

	for i := 0; i < k-1; i++ {
		node = tmpNode
		tmpNode = node.Next
		node.Next = nowNode
		nowNode = node
	}

	head.Next = ReverseKGroup(tmpNode, k)
	return nowNode
}
