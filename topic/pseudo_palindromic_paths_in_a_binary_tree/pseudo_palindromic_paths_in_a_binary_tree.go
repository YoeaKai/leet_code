package pseudo_palindromic_paths_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PseudoPalindromicPaths(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, count int) int {
	if root == nil {
		return 0
	}

	count ^= 1 << (root.Val - 1)

	res := dfs(root.Left, count) + dfs(root.Right, count)

	if (count&(count-1)) == 0 && root.Left == root.Right {
		res++
	}

	return res
}
