package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	arr := Constructor(nums)
	fmt.Println(arr.sums)
	fmt.Println(arr.SumRange(0, 5))
}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	resp := make([]int, len(nums))
	resp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		resp[i] = resp[i-1] + nums[i]
	}

	return NumArray{resp}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum := 0
	if left == 0 {
		leftSum = 0
	} else {
		leftSum = this.sums[left-1]
	}
	return this.sums[right] - leftSum
}
