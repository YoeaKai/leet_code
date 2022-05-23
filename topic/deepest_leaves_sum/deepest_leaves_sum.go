package deepest_leaves_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeepestLeavesSum(root *TreeNode) (res int) {
	var maxLevel int
	dfs(root, 0, &maxLevel, &res)
	return res
}

func dfs(node *TreeNode, curLevel int, maxLevel, sum *int) {
	if curLevel > *maxLevel {
		*maxLevel++
		*sum = node.Val
	} else if curLevel == *maxLevel {
		*sum += node.Val
	}

	if node.Left != nil {
		dfs(node.Left, curLevel+1, maxLevel, sum)
	}
	if node.Right != nil {
		dfs(node.Right, curLevel+1, maxLevel, sum)
	}
}
