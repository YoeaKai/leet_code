package unique_binary_search_trees_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) (res []*TreeNode) {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) (res []*TreeNode) {
	if start > end {
		return []*TreeNode{nil}
	}

	for i := start; i <= end; i++ {
		leftTrees, rightTrees := generate(start, i-1), generate(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}

	return res
}
