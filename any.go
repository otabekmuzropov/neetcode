package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	val := 5
	remove(nums, val)
	fmt.Println(nums)
}

func remove(nums []int, val int) {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
	}
}
