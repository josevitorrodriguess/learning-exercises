package leetcode

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	indices := make([]int, 2)

	for i, v := range nums {
		diff := target - v
		if indice, ok := table[diff]; ok {
			indices[0] = indice
			indices[1] = i
			return indices
		}
		table[v] = i
	}

	return nil
}
