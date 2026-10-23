package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 3, 1, 2}
	fmt.Println(replaceElements(arr))

}

func replaceElements(arr []int) []int {
	l := len(arr)
	// maxi := arr[l-1]
	done := make([]int, l)
	done[l-1] = -1
	for i := l - 2; i >= 0; i-- {
		done[i] = max(done[i+1], arr[i+1])
	}

	return done
}
