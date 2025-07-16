package main

import "fmt"

func main() {
	nums := []int{4, 8, 2, 10}
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] ^ nums[i]
	}

	fmt.Println(sums)

	queries := [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	for _, pairs := range queries {
		if pairs[0] == pairs[1] {
			fmt.Println(nums[pairs[0]])
		} else if pairs[0] == 0 {
			fmt.Println(sums[pairs[1]])
		} else {
			fmt.Println(sums[pairs[1]] ^ sums[pairs[0]-1])
		}
	}
}
