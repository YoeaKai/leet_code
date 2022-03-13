package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	res, v := true, root.Val
	validate(&res, root.Left, &v, nil)
	validate(&res, root.Right, nil, &v)
	return res
}

func validate(res *bool, node *TreeNode, high, low *int) {
	if node != nil && *res {
		if (high != nil && node.Val >= *high) || (low != nil && node.Val <= *low) {
			*res = false
		} else {
			v := node.Val
			validate(res, node.Left, &v, low)
			validate(res, node.Right, high, &v)
		}
	}
}
