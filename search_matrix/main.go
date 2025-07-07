package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	ROWS, COLS := len(matrix), len(matrix[0])

	top, bot := 0, ROWS-1
	var mid int

	for top <= bot {
		mid = top + (bot-top)/2

		if target > matrix[mid][COLS-1] {
			top = mid + 1
		} else if target < matrix[mid][0] {
			bot = mid - 1
		} else {
			break
		}
	}

	low, high := 0, COLS-1

	for low <= high {
		centre := low + (high-low)/2
		if target > matrix[mid][centre] {
			low = centre + 1
		} else if target < matrix[mid][centre] {
			high = centre - 1
		} else {
			return true
		}
	}

	return false
}
