package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	fmt.Println(prefixesDivBy5(nums))
}

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	resp := make([]bool, n)
	for i := range n - 1 {
		resp[i] = nums[i]%5 == 0
		nums[i+1] = (2*nums[i] + nums[i+1]) % 5
	}
	resp[n-1] = nums[n-1]%5 == 0

	return resp
}
