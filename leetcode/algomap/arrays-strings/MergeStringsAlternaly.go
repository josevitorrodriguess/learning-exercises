package leetcode

func mergeAlternately(word1 string, word2 string) string {
	var stack []string
	var wordMerged string
	i := 0

	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			stack = append(stack, string(word1[i]))
		}
		if i < len(word2) {
			stack = append(stack, string(word2[i]))
		}
		i++
	}

	for _, val := range stack {
		wordMerged += val
	}

	return wordMerged
}
