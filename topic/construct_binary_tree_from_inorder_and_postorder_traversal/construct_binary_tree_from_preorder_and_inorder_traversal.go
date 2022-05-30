package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	n := len(postorder) - 1
	idx := 0
	for i, val := range inorder {
		if val == postorder[n] {
			idx = i
			break
		}
	}

	node := &TreeNode{
		Val:   postorder[n],
		Left:  BuildTree(inorder[:idx], postorder[:idx]),
		Right: BuildTree(inorder[idx+1:], postorder[idx:n]),
	}

	return node
}
