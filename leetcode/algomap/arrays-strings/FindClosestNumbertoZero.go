package leetcode

import "math"

func findClosestNumber(nums []int) int {
	num := nums[0]

	for _, i := range nums {
		if math.Abs(float64(i)) < math.Abs(float64(num)) || (math.Abs(float64(i)) == math.Abs(float64(num)) && i > num) {
			num = i
		}
	}
	return num
}
