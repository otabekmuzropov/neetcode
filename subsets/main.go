package main

import (
	"fmt"
	"sort"
)

func backtrack(nums []int, index int, current []int, result *[][]int) {
	// Add a copy of the current subset to the result
	tmp := make([]int, len(current))
	copy(tmp, current)
	*result = append(*result, tmp)

	for i := index; i < len(nums); i++ {
		// Skip duplicates
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		// Include current element
		current = append(current, nums[i])
		backtrack(nums, i+1, current, result)
		// Backtrack
		current = current[:len(current)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // Sort to handle duplicates
	var result [][]int
	backtrack(nums, 0, []int{}, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsetsWithDup(nums))
}
