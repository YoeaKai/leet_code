package sum_of_distances_in_tree

// ref: https://leetcode.com/problems/sum-of-distances-in-tree/discuss/130583/C++JavaPython-Pre-order-and-Post-order-DFS-O(N)
func SumOfDistancesInTree(n int, edges [][]int) []int {
	tree := make(map[int][]int)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	count := make([]int, n)
	for i := 0; i < len(count); i++ {
		count[i] = 1
	}

	res := make([]int, n)

	// Find root res and pre-jobs for dfs2.
	dfs(0, -1, tree, count, res)
	dfs2(0, -1, tree, count, res)

	return res
}

// Post-order
func dfs(node, parent int, tree map[int][]int, count, res []int) {
	for _, child := range tree[node] {
		if child != parent {
			dfs(child, node, tree, count, res)
			count[node] += count[child]
			res[node] += res[child] + count[child]
		}
	}
}

// Pre-order
func dfs2(node, parent int, tree map[int][]int, count, res []int) {
	for _, child := range tree[node] {
		if child != parent {
			// Because:	res(child) - res(parent) = count(parent) - count(child)
			// and:		count(parent) = n - count(child)
			// Deduced:	res(child) - res(parent) = n - 2 count(child)
			res[child] = res[node] - count[child] + len(count) - count[child]
			dfs2(child, node, tree, count, res)
		}
	}
}
