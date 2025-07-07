package main

import "fmt"

func main() {
	nums := []int{4, -1, 2, -7, 3, 4}
	fmt.Println(kadanes(nums))
}

func kadanes(nums []int) int {
	currSum, maxSum := 0, 0
	for i := range nums {
		currSum = max(nums[i], currSum+nums[i])
		//currSum = max(currSum, 0)
		//currSum += nums[i]
		maxSum = max(maxSum, currSum)
		fmt.Println(currSum, maxSum)
	}

	return maxSum
}
