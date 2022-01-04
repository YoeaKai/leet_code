package lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findNode(root *TreeNode, p, q int) *TreeNode {
	if p > root.Val {
		if q > root.Val {
			return findNode(root.Right, p, q)
		} else {
			return root
		}
	} else if p < root.Val {
		if q < root.Val {
			return findNode(root.Left, p, q)
		} else {
			return root
		}
	} else {
		return root
	}
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findNode(root, p.Val, q.Val)
}
