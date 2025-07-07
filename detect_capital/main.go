package main

import "fmt"

func main() {
	word := "olmA"
	fmt.Println(DetectCapitalUse(word))
}

func DetectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	is1stLower := checkToLower(word[0])
	is2ndLower := checkToLower(word[1])

	for i := 1; i < len(word)-1; i++ {
		if (checkToLower(word[i]) && !checkToLower(word[i+1])) || (!checkToLower(word[i]) && checkToLower(word[i+1])) {
			return false
		}
	}

	return (is1stLower && is2ndLower) || (!is1stLower && !is2ndLower) || (!is1stLower && is2ndLower)
}

func checkToLower(b byte) bool {
	return b >= 91 && b <= 122
}
