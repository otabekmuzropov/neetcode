package main

import "fmt"

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	done := 0

	for _, num := range nums {
		if !m[num-1] {
			long := 0
			for m[num+long] {
				long++
			}
			done = max(done, long)
		}
	}

	return done
}
