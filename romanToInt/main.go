package main

import "fmt"

func main() {
	s1 := "III"
	fmt.Println(romanToInt(s1))
	s1 = "LVIII"
	fmt.Println(romanToInt(s1))
	s1 = "MCMXCIV"
	fmt.Println(romanToInt(s1))
}

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var prev, curr int
	num := 0
	prevIndex := 0
	for idx := range s {
		curr = m[s[idx]]
		prev = m[s[prevIndex]]
		if curr > prev {
			num -= 2 * prev
		}

		num += curr

		prevIndex = idx
	}

	return num
}
