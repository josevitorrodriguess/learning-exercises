package leetcode

func isSubsequence(s string, t string) bool {
	sIndex := 0
	sLen := len(s)

	for _, val := range t {
		if sIndex < sLen && s[sIndex] == byte(val) {
			sIndex++
		}
	}

	return sIndex == sLen
}
