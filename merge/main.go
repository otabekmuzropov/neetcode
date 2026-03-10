package main

import "fmt"

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	m := 0
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	for i := range n {
		nums1[i] = nums2[i]
	}
}

/*
[1,2,3,0,0,0]
[2,5,6]

1)
[1,2,3,0,0,0] - 0
[2,5,6] - 0

2)
[1,2,3,0,0,0] - 1
[2,5,6] - 0

3)
[1,2,2,0,0,0] - 2
[3,5,6] - 0

4)
[1,2,2,3,0,0] - 3
[3,5,6] - 1

5)
[1,2,2,3,5,0] - 4
[3,5,6] - 2

6)
[1,2,2,3,5,6] - 4
[3,5,6] - 2
*/
