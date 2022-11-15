package count_complete_tree_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(node *TreeNode) int {
	if node == nil {
		return -1
	}

	return 1 + height((*node).Left)
}

func CountNodes(root *TreeNode) int {
	nodes := 0
	h := height(root)

	for root != nil {
		if height(root.Right) == h-1 {
			nodes += 1 << h
			root = root.Right
		} else {
			nodes += 1 << (h - 1)
			root = root.Left
		}
		h--
	}

	return nodes
}
