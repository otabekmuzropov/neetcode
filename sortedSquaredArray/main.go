package main

import "fmt"

func main() {
	arr := []int{-7, -4, 1, 2, 3, 5, 6, 8, 9}
	fmt.Println(SortedSquaredArray(arr))
}

func SortedSquaredArray(array []int) []int {
	l := len(array)
	done := make([]int, l)
	index := l - 1

	left, right := 0, l-1
	for left <= right {
		if abs(array[left]) > abs(array[right]) {
			done[index] = array[left] * array[left]
			left++
		} else {
			done[index] = array[right] * array[right]
			right--
		}
		index--
	}

	return done
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
