package leetcode

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	number := 0
	i := 0
	n := len(s)   

	for i < n {
		if i < n-1 && m[string(s[i])] < m[string(s[i+1])] {
			number += m[string(s[i+1])] - m[string(s[i])]
			i += 2
		} else {
			number += m[string(s[i])]
			i += 1
		}

	}
	return number
}
