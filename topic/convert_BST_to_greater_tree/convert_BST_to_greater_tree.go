package convert_BST_to_greater_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConvertBST(root *TreeNode) *TreeNode {
	var sum int
	dfs(root, &sum)
	return root
}

func dfs(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	dfs(node.Right, sum)
	*sum += node.Val
	node.Val = *sum
	dfs(node.Left, sum)
}
