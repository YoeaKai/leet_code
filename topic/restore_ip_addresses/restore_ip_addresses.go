package restore_ip_addresses

import "strconv"

func RestoreIpAddresses(s string) (res []string) {
	n := len(s)

	for i := 1; i <= 3; i++ {
		for j := i + 1; j <= i+3; j++ {
			for k := j + 1; k <= j+3; k++ {
				// Front 3 part can't longer than n.
				// The forth part can't longer than 3.
				if k < n && n-k <= 3 && validate(s[:i]) && validate(s[i:j]) && validate(s[j:k]) && validate(s[k:]) {
					res = append(res, s[:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
				}
			}
		}
	}

	return res
}

func validate(s string) bool {
	if len(s) > 1 && s[0:1] == "0" {
		return false
	}

	converted, _ := strconv.Atoi(s)
	return converted <= 255
}
