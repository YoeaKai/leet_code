package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) (res []int) {
	dfs(&res, root)
	return res
}

func dfs(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		dfs(res, node.Left)
	}

	*res = append(*res, node.Val)

	if node.Right != nil {
		dfs(res, node.Right)
	}
}
