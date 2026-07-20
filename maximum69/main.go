package main

import (
	"fmt"
	"math"
)

func main() {
	num := 9669
	fmt.Println(maximum69Number(num))
}

func maximum69Number(num int) int {
	digits := []int{}
	for num != 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	l := len(digits)
	found := false
	res := 0

	for i := l - 1; i >= 0; i-- {
		ith := digits[i]
		if ith == 6 && !found {
			ith = 9
			found = true
		}

		res += ith * int(math.Pow10(i))
	}

	return res
}
