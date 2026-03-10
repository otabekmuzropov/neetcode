package main

import "fmt"

func main() {
	nums := []int{3, 9, 7}
	k := 5
	fmt.Println(minOperations(nums, k))
}

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum % k
}
