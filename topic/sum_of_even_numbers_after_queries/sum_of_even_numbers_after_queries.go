package sum_of_even_numbers_after_queries

func SumEvenAfterQueries(nums []int, queries [][]int) []int {
	result := make([]int, len(queries))

	var even int
	for _, num := range nums {
		if num%2 == 0 {
			even += num
		}
	}

	for i := 0; i < len(queries); i++ {
		if nums[queries[i][1]]%2 == 0 {
			if queries[i][0]%2 == 0 {
				even += queries[i][0]
			} else {
				even -= nums[queries[i][1]]
			}
		} else if queries[i][0]%2 != 0 { // nums[queries[i][1]] and value of queries are odd.
			even += (nums[queries[i][1]] + queries[i][0])
		}

		result[i] = even
		nums[queries[i][1]] += queries[i][0]
	}

	return result
}
