// Input two arrays A and B of N integers each, describing a directed graph,
// returns true if the graph is a cycle and false otherwise.
// e.g. A = [1, 3, 2, 4] and B = [4, 1, 3, 2]
// Output: true

package interview_PixlS_2022_11

func Solution2(A []int, B []int) bool {
	graph := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := graph[A[i]]; ok {
			return false
		}
		graph[A[i]] = B[i]
	}

	count := 1
	curr := B[0]

	for curr != A[0] {
		curr = graph[curr]
		count++
	}

	return count == len(A)
}
