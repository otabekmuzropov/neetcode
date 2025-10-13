package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))

}

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	short := l + 1
	total := 0
	L := 0

	for R := range l {
		total += nums[R]
		for total >= target {
			short = min(short, R-L+1)
			total -= nums[L]
			L++
		}
		R++
	}

	if short == l+1 {
		return 0
	}

	return short
}
