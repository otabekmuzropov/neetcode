package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 7, 3, 5, -1}
	fmt.Println(bucketSort(arr)) // 0,0,1,1,1,2,2
	// fmt.Println(bucket(arr))
}

func bucket(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var maxi, mini = arr[0], arr[0]
	for _, val := range arr {
		if val > maxi {
			maxi = val
		}

		if val < mini {
			mini = val
		}
	}

	intended := make([]int, maxi+1)

	for _, val := range arr {
		intended[val]++
	}

	idx := 0
	for i := 1; i < len(intended); i++ {
		ith := intended[i]
		for range ith {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(intended)

	return arr
}

func bucketSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}

	maxi, mini := arr[0], arr[0]
	for _, val := range arr {
		maxi = max(maxi, val)
		mini = min(mini, val)
	}

	shift := 0

	if mini < 0 {
		shift = -mini
	}

	out := make([]int, maxi+shift+1)

	for _, val := range arr {
		out[val+shift]++
	}

	i := 0
	for idx, val := range out {
		for range val {
			arr[i] = idx - shift
			i++
		}
	}

	return arr
}
