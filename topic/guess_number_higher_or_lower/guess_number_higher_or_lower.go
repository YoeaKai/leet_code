package guess_number_higher_or_lower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int

func GuessNumber(n int) int {
	max := n
	min := 1

	for {
		mid := (max-min)/2 + min
		num := guess(mid)
		if num == 1 {
			min = mid + 1
		} else if num == -1 {
			max = mid - 1
		} else {
			return mid
		}
	}
}
