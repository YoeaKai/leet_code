package range_sum_of_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return RangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return RangeSumBST(root.Right, low, high)
	}

	return root.Val + RangeSumBST(root.Left, low, high) + RangeSumBST(root.Right, low, high)
}
