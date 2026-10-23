package main

import (
	"fmt"
)

func main() {
	s := "node"
	t := "neetcode"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls > lt {
		return false
	}

	k := 0

	for i := 0; i < lt; i++ {
		fmt.Println(i)
		if s[k] == t[i] {
			k++
		}
	}

	fmt.Println(k)

	return k == ls
}
