package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(node *TreeNode, cur int) int {
	if node == nil {
		return 0
	}

	if node.Right == nil && node.Left == nil {
		return cur*10 + node.Val
	}

	return sum(node.Left, cur*10+node.Val) + sum(node.Right, cur*10+node.Val)
}
