package main

import "fmt"

func main() {
	nums := []int{10, 4, 8, 3}
	result := LeftRightDifference(nums)
	fmt.Println(result)
}

func LeftRightDifference(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	done := make([]int, l)
	left[0] = 0
	right[l-1] = 0
	for i := 1; i < l; i++ {
		left[i] = left[i-1] + nums[i-1]
		right[l-i-1] = right[l-i] + nums[l-i]
	}

	for i := 0; i < l; i++ {
		diff := left[i] - right[i]
		switch {
		case diff > 0:
			done[i] = -1
		case diff < 0:
			done[i] = 1
		default:
			done[i] = 0
		}
	}

	return done
}
