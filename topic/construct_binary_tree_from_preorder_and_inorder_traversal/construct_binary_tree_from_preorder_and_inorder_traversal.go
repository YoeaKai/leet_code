package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	idx := 0
	for i, val := range inorder {
		if val == preorder[0] {
			idx = i
			break
		}
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:idx+1], inorder[:idx]),
		Right: BuildTree(preorder[idx+1:], inorder[idx+1:]),
	}

	return root
}
