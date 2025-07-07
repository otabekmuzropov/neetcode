package main

import "fmt"

func main() {
	fmt.Println(Generate(5))
}

func Generate(n int) [][]int {
	resp := make([][]int, n)
	for i := 0; i < n; i++ {
		nth := make([]int, i+1)
		nth[0] = 1
		nth[len(nth)-1] = 1
		for j := 1; j < i; j++ {
			nth[j] = resp[i-1][j-1] + resp[i-1][j]
		}

		resp[i] = nth
	}

	return resp
}
