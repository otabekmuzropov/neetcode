package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	done := make([][]int, 0)

	slices.Sort(nums)

	for idx, num := range nums {
		if idx > 0 && num == nums[idx-1] {
			continue
		}

		l, r := idx+1, len(nums)-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				done = append(done, []int{num, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return done
}
