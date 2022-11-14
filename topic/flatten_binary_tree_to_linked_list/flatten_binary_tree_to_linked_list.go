package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1
func Flatten(root *TreeNode) {
	flattenTree(root, nil)
}

func flattenTree(root, pre *TreeNode) *TreeNode {
	if root == nil {
		if pre != nil {
			return pre
		}
		return nil
	}

	right := flattenTree(root.Right, pre)
	left := flattenTree(root.Left, right)
	root.Right = left
	root.Left = nil

	return root
}

// Method 2: with global variable.
var pre2 *TreeNode

func Flatten2(root *TreeNode) {
	pre2 = nil
	flattenTree2(root)
}

func flattenTree2(root *TreeNode) {
	if root == nil {
		return
	}

	flattenTree2(root.Right)
	flattenTree2(root.Left)
	root.Right = pre2
	root.Left = nil
	pre2 = root
}
