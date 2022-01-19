package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	res := make([]int, len(nums))
	midum := (len(nums) + 1) / 2
	res[0] = 1
	res[len(nums)-1] = 1

	for i := 1; i < midum; i++ {
		j := len(nums) - 1 - i
		res[i] = res[i-1] * nums[i-1]
		res[j] = res[j+1] * nums[j+1]
	}

	left := res[midum-2] * nums[midum-2]
	right := res[midum]

	if len(nums)%2 == 1 {
		right = res[midum-1]
		res[midum-1] = right * left
	}

	for i := midum; i < len(nums); i++ {
		j := len(nums) - 1 - i
		left *= nums[i-1]
		right *= nums[j+1]
		res[i] *= left
		res[j] *= right
	}

	return res
}

func ProductExceptSelfII(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		res[i] *= tmp
	}

	return res
}
