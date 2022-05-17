package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return isSymmetricAlgo(root.Left, root.Right)
}

func isSymmetricAlgo(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricAlgo(left.Left, right.Right) && isSymmetricAlgo(left.Right, right.Left)
}
