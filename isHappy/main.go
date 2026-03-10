package main

import "fmt"

func main() {
	n := 2
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true

	for n != 1 {
		str := fmt.Sprintf("%d", n)
		sum := 0
		for _, s := range str {
			a := int(s - '0')
			sum += a * a
		}

		n = sum
		if m[n] {
			return false
		}

		m[n] = true
	}

	return true
}
