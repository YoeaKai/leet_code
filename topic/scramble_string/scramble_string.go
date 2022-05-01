package scramble_string

// Method 1: Tabulation
func IsScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]bool, n)

	// Initialization.
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}

	// dp[i][j][k]:
	// i : means start index of s1
	// j : means start index of s2
	// k : means the length starts with i and j in s1 and s2.
	for k := 1; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			for j := 0; j+k <= n; j++ {
				if k == 1 {
					dp[i][j][k] = s1[i] == s2[j]
				} else {
					for q := 1; q < k && !dp[i][j][k]; q++ {
						// Split into two part: (0 to q) and (q+1 to k) for s1 and s2.
						// Stop the loop if swap or not swap has true result.
						dp[i][j][k] = (dp[i][j][q] && dp[i+q][j+q][k-q]) || (dp[i][j+k-q][q] && dp[i+q][j][k-q])
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

// Method 2: Memoization (Result: TLE)
func IsScramble2(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) != len(s2) {
		return false
	}

	// Compare letter
	letters := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		letters[int(s1[i])-int('a')]++
		letters[int(s2[i])-int('a')]--
	}
	for i := 0; i < len(s1); i++ {
		if letters[int(s1[i])-int('a')] != 0 {
			return false
		}
	}

	for i := 1; i < len(s1); i++ {
		if IsScramble2(s1[0:i], s2[0:i]) && IsScramble2(s1[i:], s2[i:]) {
			return true
		}
		if IsScramble2(s1[0:i], s2[len(s2)-i:]) && IsScramble2(s1[i:], s2[0:len(s2)-i]) {
			return true
		}
	}

	return false
}
