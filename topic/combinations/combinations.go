package combinations

// C: The combinations of taking k values in n.
// C = P / k! = n! / (k! * (n - k)!)
func Combine(n int, k int) [][]int {
	res := [][]int{}
	dfs(&res, []int{}, n, k)
	return res
}

func dfs(res *[][]int, store []int, n int, k int) {
	if len(store) == k {
		tmp := make([]int, len(store))
		copy(tmp, store)
		*res = append(*res, tmp)
	} else {
		for i := n; i > 0; i-- {
			store = append(store, i)
			dfs(res, store, i-1, k)
			store = store[:len(store)-1]
		}
	}
}
