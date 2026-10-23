package main

import "fmt"

func main() {
	n := 19
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		m[n] = true
		n = sumOfSqr(n)
		if m[n] {
			return false
		}
	}

	return true
}

func sumOfSqr(n int) int {
	sum := 0
	for n != 0 {
		reminder := n % 10
		n /= 10
		sum += reminder * reminder
	}

	return sum
}

func myPow(x float64, n int) float64 {
	return 1.0000
}