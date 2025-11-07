package main

import "fmt"

func main() {
	a := 10
	fmt.Println(smallestNumber(a))
}

func smallestNumber(n int) int {
	a := 0
	for n != 0 {
		n >>= 1
		a++
	}

	return 1<<a - 1
}
