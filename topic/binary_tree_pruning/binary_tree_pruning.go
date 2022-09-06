package binary_tree_pruning

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
