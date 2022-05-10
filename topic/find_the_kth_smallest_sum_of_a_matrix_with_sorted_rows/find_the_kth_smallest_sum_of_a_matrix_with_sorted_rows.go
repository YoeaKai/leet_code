package find_the_kth_smallest_sum_of_a_matrix_with_sorted_rows

import "container/heap"

type Value struct {
	num1   int // The value of nums1.
	index2 int // The index of nums2.
	value  int // The value of num1 + nums2[index2].
}

type Queue []*Value

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(a, b int) bool {
	return q[a].value < q[b].value
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*Value))
}

func (q *Queue) Pop() interface{} {
	n := len(*q) - 1
	min := (*q)[n]
	*q = (*q)[:n]
	return min
}

func KthSmallest(mat [][]int, k int) int {
	res := mat[0]

	for _, row := range mat[1:] {
		res = kSmallestPairs(res, row, k)
	}

	return res[k-1]
}

func kSmallestPairs(nums1, nums2 []int, k int) (res []int) {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return res
	}

	var queue Queue

	// Push all value of nums1 add nums2[0].
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&queue, &Value{
			num1:   nums1[i],
			index2: 0,
			value:  nums1[i] + nums2[0],
		})
	}

	// Record nums2 corresponding to the value of nums1, and iterate the index of nums2.
	// The length of res is limited to k.
	for ; k != 0 && len(queue) != 0; k-- {
		cur := heap.Pop(&queue).(*Value)
		res = append(res, cur.value)

		if cur.index2 == len(nums2)-1 {
			continue
		}

		heap.Push(&queue, &Value{
			num1:   cur.num1,
			index2: cur.index2 + 1,
			value:  cur.num1 + nums2[cur.index2+1],
		})
	}

	return res
}
