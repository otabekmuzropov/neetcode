package main

import "fmt"

func main() {
	fmt.Println(climbStairs(8))
}

func fibonacci(n int) int {
	sequence := make([]int, n+1)
	sequence[0] = 0
	sequence[1] = 1
	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	fmt.Println(sequence)
	return sequence[n]
}

func climbStairs(n int) int {
	return fibonacci(n + 1)
}
