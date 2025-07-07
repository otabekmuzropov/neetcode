package main

import "fmt"

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	resp := make([]int, len(nums))
	counter := 0

	for i < j {
		startSqr := nums[i] * nums[i]
		endSqr := nums[j] * nums[j]
		if startSqr > endSqr {
			fmt.Println("numsi", nums[i])
			resp[len(resp)-1-counter] = startSqr
			counter++
			i++
		} else {
			fmt.Println("numsj", nums[j])

			resp[len(resp)-1-counter] = endSqr
			counter++
			j--
		}
	}

	resp[len(resp)-1-counter] = nums[i] * nums[j]

	return resp
}
