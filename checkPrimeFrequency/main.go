package main

import (
	"fmt"
	"runtime"
)

func main() {
	nums := []int{3, 0, 3, 6, 3, 3}
	fmt.Println(checkPrimeFrequency(nums))
}

func checkPrimeFrequency(nums []int) bool {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for _, val := range m {
		if val >= 2 && isPrime(val) {
			runtime.GC()
			return true
		}
	}

	runtime.GC()

	return false
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
