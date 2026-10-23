package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {

	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}

	fmt.Println(nums)
	return len(nums)
}
