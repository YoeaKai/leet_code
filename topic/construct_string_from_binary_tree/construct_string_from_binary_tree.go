package construct_string_from_binary_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Tree2str(root *TreeNode) (result string) {
	if root == nil {
		return result
	}

	result += fmt.Sprint(root.Val)

	if root.Left == nil && root.Right == nil {
		return result
	}

	result += "(" + Tree2str(root.Left) + ")"

	if root.Right != nil {
		result += "(" + Tree2str(root.Right) + ")"
	}

	return result
}
