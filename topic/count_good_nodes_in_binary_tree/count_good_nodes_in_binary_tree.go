package count_good_nodes_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GoodNodes(root *TreeNode) int {
	return dfs(root, -10000)
}

func dfs(root *TreeNode, prev int) (count int) {
	if root == nil {
		return 0
	}

	if root.Val >= prev {
		count++
	}

	return count + dfs(root.Left, max(prev, root.Val)) + dfs(root.Right, max(prev, root.Val))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
