package main

import "fmt"

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{25}
	fmt.Println(IsValidSubsequence(array, sequence))

}

func IsValidSubsequence(array []int, sequence []int) bool {
	// yechimingizni yozing.
	i := 0
	j := 0

	for j < len(array) {
		if array[j] == sequence[i] {
			i++
		}
		j++

		if i == len(sequence) {
			return true
		}
	}

	return false
}
