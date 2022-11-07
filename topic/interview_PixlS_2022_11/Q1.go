// Input an array representing the features of cars.
// Find an array of integers denoting,
// for every car in cars, the number of other similar cars.
// Two cars are similar if their descriptions differ by at most one feature.
// For example "01101" and "01001" are similar because they differ
// only by feature number 2.
// e.g. cars = ["100", "110", "010", "011", "100"]
// Output: [2, 3, 2, 1, 2]

package interview_PixlS_2022_11

func Solution1(cars []string) []int {
	length := len(cars)
	res := make([]int, length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if isSimilar(cars[i], cars[j]) {
				res[i]++
				res[j]++
			}
		}
	}

	return res
}

func isSimilar(s1, s2 string) bool {
	differ := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			differ++
		}
		if differ > 1 {
			return false
		}
	}

	return true
}
