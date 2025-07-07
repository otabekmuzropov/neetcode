package main

import (
	"fmt"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, str := range strs {
		iterated := [26]byte{}
		for _, val := range str {
			iterated[val-97]++
		}

		m[iterated] = append(m[iterated], str)
	}

	resp := make([][]string, len(m))
	i := 0
	for _, val := range m {
		resp[i] = val
		i++
	}

	return resp
}
