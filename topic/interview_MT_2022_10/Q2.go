// Find the max sum of the substring with consecutive numbers.
// e.g. A = [1,3,6,1,6,6,9,9]
// Output: 19 ([6,1,6,6])

package interview_MT_2022_10

func Q2(A []int) (res int) {
	numPos := make(map[int]int)
	sum := make([]int, len(A))
	numPos[A[0]] = 0
	sum[0] = A[0]

	for i := 1; i < len(A); i++ {
		sum[i] = sum[i-1] + A[i]
		if _, ok := numPos[A[i]]; ok {
			res = max(res, sum[i]-sum[numPos[A[i]]]+A[i])
		} else {
			numPos[A[i]] = i
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Sliding window method
func Q2_2(A []int) int {
	length := len(A)
	nums := make(map[int]int)

	for i := 0; i < length; i++ {
		nums[A[i]]++
	}

	head := -1

	for i := 0; i < length; i++ {
		if nums[A[i]] >= 2 {
			head = i
			break
		}
	}

	if head == -1 {
		return -1
	}

	var maxSum int
	var sum int
	var tail int

	for tail = head; nums[A[head]] > 0; tail++ {
		sum += A[tail]
		nums[A[tail]]--
	}

	tail--
	maxSum = sum

	for i := head + 1; i < length; i++ {
		sum -= A[i-1]
		if nums[A[i]] != 0 {
			for j := tail + 1; nums[A[i]] > 0; j++ {
				sum += A[j]
				nums[A[j]]--
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
