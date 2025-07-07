package main

import (
	"fmt"
)

func main() {
	words := "k"
	fmt.Println(LongestPalindrome(words))
}

func LongestPalindrome(s string) string {
	resp := ""
	maxV := -1
	lngth := len(s)
	for i := 0; i < lngth; i++ {
		for j := lngth - 1; j >= i; j-- {
			fmt.Println("here again", s[i:j+1], j-i)
			if isPalindrome(s[i:j+1]) && maxV < j-i {
				resp = s[i : j+1]
				maxV = j - i
			}
		}
	}

	return resp
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
