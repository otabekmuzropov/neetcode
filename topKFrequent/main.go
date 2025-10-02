package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{}, 1))
}

func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	return nums
}
