package nonnegative_integers_without_consecutive_ones

func FindIntegers(n int) int {
	if n == 0 {
		return 0
	}

	cnt := 0

	for ; 1<<cnt <= n; cnt++ {
	}

	DP := make([]int, cnt+1)
	DP[0] = 1
	DP[1] = 1

	for i := 2; i <= cnt; i++ {
		DP[i] = DP[i-1] + DP[i-2]
	}

	last := false
	thisNumIs := true
	result := 0

	for cnt -= 1; cnt >= 0; cnt-- {
		if 1<<cnt&n == 1<<cnt { //curr postion is 1
			result += DP[cnt+1]
			if last {
				thisNumIs = false
				break
			}
			last = true
		} else { //curr position is 0
			last = false
		}
	}

	if thisNumIs {
		result += 1
	}

	return result
}
