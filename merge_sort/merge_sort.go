package main

import "fmt"

func main() {
	arr := []int{32, 134, 2, 3131, 23, 134, 234, 23, 43, 423, 4}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	resp := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if right[j] < left[i] {
			resp = append(resp, right[j])
			j++
		} else {
			resp = append(resp, left[i])
			i++
		}
	}

	resp = append(resp, left[i:]...)
	resp = append(resp, right[j:]...)

	return resp
}
