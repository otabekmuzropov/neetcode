package main

import (
	"fmt"
	"neetcode/algo-io/bitwise"
)

func main() {
	pos := 1
	n := 10

	fmt.Println(6 & -6)
	fmt.Println(bitwise.ClearPosition(n, pos))
	// 11101
	//  1010
}
