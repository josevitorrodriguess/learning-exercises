package leetcode

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	indices := make([]int, 2)

	for i, v := range nums {
		diff := target - v
		if indice, ok := hashmap[diff]; ok {
			indices[0] = indice
			indices[1] = i
			return indices
		}
		hashmap[v] = i
	}
 
	return indices
}
