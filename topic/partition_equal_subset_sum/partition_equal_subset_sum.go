package partition_equal_subset_sum

func CanPartition(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2 // One part of the sum that can make the total sum.

	if nums[0] == sum {
		return true
	}

	dp := make([]bool, sum+1)
	dp[0] = true

	for _, num := range nums {
		for i := sum; i > 0; i-- {
			if i >= num {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}

	return dp[sum]
}

func CanPartition2(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2 // One part of the sum that can make the total sum.

	if nums[0] == sum {
		return true
	}

	p := make(map[int]bool)
	p[nums[0]] = true

	for i := 1; i < len(nums); i++ {
		tmp := []int{}

		for key := range p {
			added := key + nums[i]

			if added == sum {
				return true
			}

			_, ok := p[added]

			if !ok && added < sum {
				tmp = append(tmp, added)
			}
		}

		for _, val := range tmp {
			p[val] = true
		}
	}

	return false
}
