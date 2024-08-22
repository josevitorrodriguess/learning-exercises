package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {

	prefix := strs[0]
	for _, str := range strs[1:] {
		for !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
