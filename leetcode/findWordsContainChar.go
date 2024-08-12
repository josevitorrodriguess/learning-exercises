package leetcode

import "strings"

func findWordsContaining(words []string, x byte) []int {
	var idx []int

	for i := 0; i < len(words); i++ {

		if strings.Contains(words[i], string(x)) {
			idx = append(idx, i)
		}
	}
	return idx
}
