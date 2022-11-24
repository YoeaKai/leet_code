package binary_tree_maximum_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func MaxPathSum(root *TreeNode) int {
	res := MinInt
	sum(root, &res)
	return res
}

func sum(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}

	left := max(sum(node.Left, res), 0)
	right := max(sum(node.Right, res), 0)
	(*res) = max((*res), left+right+node.Val)

	return max(left, right) + node.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
