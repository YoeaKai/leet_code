package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) (res [][]int) {
	travel(root, &res, 0)
	return res
}

// DFS
func travel(node *TreeNode, res *[][]int, level int) {
	if node == nil {
		return
	}

	if len(*res) <= level {
		*res = append(*res, []int{})
	}

	if level%2 == 0 {
		(*res)[level] = append((*res)[level], node.Val)
	} else {
		(*res)[level] = append([]int{node.Val}, (*res)[level]...)
	}

	travel(node.Left, res, level+1)
	travel(node.Right, res, level+1)
}
