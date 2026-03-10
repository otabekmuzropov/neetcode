package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 4, 5, 7, 9}
	fmt.Println(summaryRanges(nums))

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))

	nums = []int{2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	resp := make([]string, 0)
	if len(nums) == 0 {
		return resp
	}
	pairs := make([][2]int, 0)
	pair := [2]int{}
	pair[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			pair[1] = nums[i-1]
			pairs = append(pairs, pair)
			pair = [2]int{nums[i], 0}
		}
	}

	pair[1] = nums[len(nums)-1]
	pairs = append(pairs, pair)

	for _, pair := range pairs {
		str := ""
		if pair[0] == pair[1] {
			str = fmt.Sprintf("%d", pair[0])
		} else {
			str = fmt.Sprintf("%d->%d", pair[0], pair[1])
		}

		resp = append(resp, str)

	}

	return resp
}
