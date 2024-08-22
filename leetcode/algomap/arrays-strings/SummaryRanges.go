package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	var arr []string
	n := len(nums)
	i := 0

	for i < n {
		start := i
		for i+1 < n && nums[i]+1 == nums[i+1] {
			i++
		}
		if start == i {
			arr = append(arr, strconv.Itoa(nums[start]))
		} else {
			arr = append(arr, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i]))
		}
		i++
	}

	return arr
}
