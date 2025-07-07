package main

import "fmt"

func main() {
	nums1 := []int{2, 3, 4}
	nums2 := []int{1, 6}
	fmt.Println(mergeArrays(nums1, nums2))
}

func mergeArrays(nums1, nums2 []int) []int {
	var resp []int
	// idx1, idx2 := len(nums1), len(nums2)
	idx1, idx2 := 0, 0

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			resp = append(resp, nums1[idx1])
			idx1++
		} else {
			resp = append(resp, nums2[idx2])
			idx2++
		}
	}

	resp = append(resp, nums1[idx1:]...)
	resp = append(resp, nums2[idx2:]...)

	return resp
}
