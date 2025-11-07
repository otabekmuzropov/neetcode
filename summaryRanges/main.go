package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
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
	str := fmt.Sprintf("%d", nums[0])
	pairs := make([][2]int, 0)
	pair := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {

		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			str = fmt.Sprintf("%s->%d", str, nums[i-1])
			resp = append(resp, str)
			// if i != len(nums)-1 {
			// 	str = fmt.Sprintf("%d", nums[i+1])
			// 	i++
			// }
			str = ""
		} else {

		}
	}

	return resp
}
