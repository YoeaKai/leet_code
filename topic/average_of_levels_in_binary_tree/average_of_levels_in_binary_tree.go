package average_of_levels_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) (result []float64) {
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n, sum := len(queue), 0

		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}

			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}

			sum += pop.Val
		}

		result = append(result, float64(sum)/float64(n))
	}

	return result
}
