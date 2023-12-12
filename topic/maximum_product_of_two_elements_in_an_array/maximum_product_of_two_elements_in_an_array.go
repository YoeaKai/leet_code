package maximum_product_of_two_elements_in_an_array

func MaxProduct(nums []int) int {
	m1 := 0
	m2 := 0

	for _, num := range nums {
		if num > m1 {
			m2 = m1
			m1 = num
		} else if num > m2 {
			m2 = num
		}
	}

	return (m1 - 1) * (m2 - 1)
}
