package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	suffix := make([]int, l)

	prefix[0], prefix[1] = 1, nums[0]
	suffix[0], suffix[1] = 1, nums[l-1]

	for i := 2; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		suffix[i] = suffix[i-1] * nums[l-i]
	}

	for i := range l {
		nums[i] = prefix[i] * suffix[l-i-1]
	}

	return nums
}
