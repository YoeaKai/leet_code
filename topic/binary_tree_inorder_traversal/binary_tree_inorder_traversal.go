package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1
func InorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		// push
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, root.Val)
		root = root.Right
	}

	return res
}

// Method 2
func InorderTraversal2(root *TreeNode) (res []int) {
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
