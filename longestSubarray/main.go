package main

import "fmt"

func main() {
	nums := []int{4, 2, 2, 3, 3, 3}
	fmt.Println(longestSubarray(nums))
}

func longestSubarray(nums []int) int {
	done := 1
	L := 0
	for R := range nums {
		if nums[L] != nums[R] {
			L = R
		}

		done = max(done, R-L+1)
	}

	return done
}
