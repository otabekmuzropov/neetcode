package main

import "fmt"

func main() {
	nums1 := []int{3, 0, 1}
	fmt.Println(missingNumber(nums1))
	// nums2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	// fmt.Println(missingNumber(nums2))

}

func missingNumber(nums []int) int {
	xor := nums[0]
	l := len(nums)
	for i := 1; i < l; i++ {
		xor ^= nums[i]
	}

	for i := range l + 1 {
		xor ^= i
	}

	return xor
}
