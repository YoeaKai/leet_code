package convert_sorted_list_to_binary_search_tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return toBST(head, nil)
}

func toBST(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	mid := head
	stopPos := head

	for stopPos != tail && stopPos.Next != tail {
		mid = mid.Next
		stopPos = stopPos.Next.Next
	}

	midNode := &TreeNode{
		Val:   mid.Val,
		Left:  toBST(head, mid),
		Right: toBST(mid.Next, tail),
	}

	return midNode
}
