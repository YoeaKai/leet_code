package recover_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func RecoverTree(root *TreeNode) {
	var (
		pre        = &TreeNode{Val: MinInt}
		firstNode  *TreeNode
		secondNode *TreeNode
	)

	dfs(root, &pre, &firstNode, &secondNode)

	if firstNode != nil && secondNode != nil {
		firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
	}
}

func dfs(node *TreeNode, pre, firstNode, secondNode **TreeNode) {
	if node.Left != nil {
		dfs(node.Left, pre, firstNode, secondNode)
	}

	if node.Val < (*pre).Val {
		if (*firstNode) == nil {
			*firstNode = *pre
		}
		if (*firstNode) != nil {
			*secondNode = node
		}
	}

	*pre = node

	if node.Right != nil {
		dfs(node.Right, pre, firstNode, secondNode)
	}
}
