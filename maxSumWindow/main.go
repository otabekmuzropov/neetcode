package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 1, 3, 2, 4, 9}
	k := 3
	fmt.Println(maxSumWindow(nums, k))
}

func maxSumWindow(nums []int, k int) int {
	best := 0
	l := len(nums)
	for i := range k {
		best += nums[i]
	}

	fmt.Println(best)

	curr := best

	for i := 1; i <= l-k; i++ {
		curr = curr + nums[i+k-1] - nums[i-1]
		best = max(best, curr)
	}

	return best
}
