package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4}
	fmt.Println(minimumLength(nums, 6))
}

func minimumLength(nums []int, target int) int {
	done := len(nums)
	sum := 0

	L := 0
	for R := range nums {
		sum += nums[R]
		for sum >= target {
			done = min(done, R-L+1)
			sum -= nums[L]
			L++
		}

	}

	return done
}
