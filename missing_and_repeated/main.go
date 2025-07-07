package main

import "fmt"

func main() {
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(findMissingAndRepeatedValues(grid))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	ans := make([]int, 2)
	n := len(grid) * len(grid)
	summ := n * (n + 1) / 2
	m := make(map[int]bool)
	for i := range grid {
		for j := range grid[i] {
			if m[grid[i][j]] {
				ans[0] = grid[i][j]
			} else {
				summ -= grid[i][j]
				m[grid[i][j]] = true
			}
		}
	}
	ans[1] = summ

	return ans
}
