package main

import "fmt"

func main() {
	arr := []int{4, 1, 2, 2, 3, 3, 3}
	fmt.Println(longest(arr))

	nums := []int{1, 2, 3, 4, 5, 6, 30}
	fmt.Println(shortest(nums, 100))
}

func longest(arr []int) int {
	long := 1

	l := 0
	for r := 1; r < len(arr); r++ {
		if arr[l] != arr[r] {
			l = r
		}

		long = max(long, r-l+1)
	}

	return long
}

func shortest(arr []int, target int) int {
	l := len(arr)

	sh := l
	total := 0
	L := 0

	for R := range l {
		total += arr[R]
		for total >= target {
			sh = min(sh, R-L+1)
			total -= arr[L]
			L++
		}

		R++
	}

	if sh == l {
		return 0
	}

	return sh
}
