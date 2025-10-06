package main

import "fmt"

func main() {
	nums := []int{}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	return countSort(nums)
}

func countSort(arr []int) []int {
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
