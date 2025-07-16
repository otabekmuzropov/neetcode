package main

import (
	"fmt"
)

func FindUnique(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}

func main() {
	fmt.Println(12 >> 0)

}

func replace(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}

func countBits(x int) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

func countBitsN(x int) []int {
	resp := make([]int, x+1)
	for i := 0; i <= x; i++ {
		resp[i] = countBits(i)
	}

	return resp
}

func countBitsN2(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

func IsPowerOfFour(n int) bool {
	for n != 1 {
		if n != n>>2<<2 {
			return false
		}
		n = n >> 2
	}

	return true
}
