package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		m1[s[i]]++
		m2[t[i]]++
	}

	for k, val := range m1 {
		if m2[k] != val {
			return false
		}
	}

	return true
}
