package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode

func Flatten(root *TreeNode) {
	pre = nil
	flattenTree(root)
}

func flattenTree(root *TreeNode) {
	if root == nil {
		return
	}

	flattenTree(root.Right)
	flattenTree(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}
