package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{1, 2, 3, 4, 5, 6, 7}
	// coins := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println(NonConstructibleChange(coins))

}
func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	var currentChange int
	for _, coin := range coins {
		if coin > currentChange+1 {
			return currentChange + 1
		} else {
			currentChange += coin
		}
	}

	return currentChange + 1
}
