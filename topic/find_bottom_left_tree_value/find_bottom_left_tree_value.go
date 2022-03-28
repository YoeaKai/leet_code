package find_bottom_left_tree_value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindBottomLeftValue(root *TreeNode) (res int) {
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}

		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}

		res = queue[0].Val
		queue = queue[1:]
	}

	return res
}
