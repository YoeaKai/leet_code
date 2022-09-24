package path_sum_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) (result [][]int) {
	var path []int
	dfs(root, targetSum, &path, &result)
	return result
}

func dfs(node *TreeNode, targetSum int, path *[]int, result *[][]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		*result = append(*result, append([]int(nil), append(*path, node.Val)...))
		return
	}

	*path = append(*path, node.Val)

	dfs(node.Left, targetSum-node.Val, path, result)
	dfs(node.Right, targetSum-node.Val, path, result)

	*path = (*path)[:len(*path)-1]
}
