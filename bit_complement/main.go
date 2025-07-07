package main

import "fmt"

func main() {
	fmt.Println(complement(9))
}

func complement(n int) int {
	var count int
	x := n
	for n > 1 {
		n >>= 1
		count++
	}

	resp := 1<<(count+1) - 1
	return x ^ resp
}
