package main

import (
	"fmt"
	"runtime"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	count := 0
	runtime.GC()

	for i := range nums {
		if nums[i] == 0 {
			count++
			continue
		}

		nums[i], nums[i-count] = nums[i-count], nums[i]
	}
}
