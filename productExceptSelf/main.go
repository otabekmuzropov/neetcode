package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
	nums1 := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf1(nums1))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	suffix := make([]int, l)
	prefix[0] = nums[0]
	suffix[0] = nums[l-1]

	for i := 1; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i]
		suffix[i] = suffix[i-1] * nums[l-i-1]
	}

	for i := range l {
		nums[i] = prefix[i] * suffix[l-i-1]
	}

	fmt.Println(prefix, suffix, nums)

	return []int{}

}

func productExceptSelf1(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	suffix := make([]int, l)

	prefix[0], prefix[1] = 1, nums[0]
	suffix[0], suffix[1] = 1, nums[l-1]

	for i := 2; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		suffix[i] = suffix[i-1] * nums[l-i]
	}

	for i := 0; i < l; i++ {
		nums[i] = prefix[i] * suffix[l-i-1]
	}

	return nums
}
