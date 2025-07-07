package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 5, 6, 9}
	target := 9

	fmt.Println(binarySearch(arr, target))

}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
