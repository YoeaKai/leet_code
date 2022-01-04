package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []TreeNode{*root}
	tmpQueue := []TreeNode{}
	tmp := []int{}
	for len(queue) != 0 || len(tmpQueue) != 0 {
		if len(queue) == 0 {
			queue = tmpQueue
			result = append(result, tmp)
			tmpQueue = nil
			tmp = nil
		} else {
			tmp = append(tmp, queue[0].Val)
			if queue[0].Left != nil {
				tmpQueue = append(tmpQueue, *queue[0].Left)
			}
			if queue[0].Right != nil {
				tmpQueue = append(tmpQueue, *queue[0].Right)
			}
			queue = queue[1:]
		}
	}

	if tmp != nil {
		result = append(result, tmp)
	}

	return result
}
