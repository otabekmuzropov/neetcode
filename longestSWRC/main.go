package main

import "fmt"

func main() {
	s := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	L := 0

	if len(s) == 0 {
		return 0
	}

	done := 1
	m := map[byte]int{
		s[0]: 0,
	}

	for R := 1; R < len(s); R++ {
		if val, ok := m[s[R]]; ok {
			L = val + 1
		} else {
			done = max(done, R-L+1)
		}
		m[s[R]] = R
	}

	return done
}
