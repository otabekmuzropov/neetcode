package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //pwwkew
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    //pwwkew
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   //pwwkew
	fmt.Println(lengthOfLongestSubstring("aab"))      //pwwkew
}

func lengthOfLongestSubstring(s string) int {
	L := 0
	long := 0

	m := make(map[byte]bool)

	for R := 0; R < len(s); {
		for m[s[R]] {
			delete(m, s[L])
			L++
		}

		long = max(long, R-L+1)
		m[s[R]] = true
		R++

	}

	return long
}
