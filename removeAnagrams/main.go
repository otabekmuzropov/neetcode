package main

import "fmt"

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	fmt.Println(removeAnagrams(words))
}

func removeAnagrams(words []string) []string {
	for i := 1; i < len(words); {
		if isAnagram(words[i], words[i-1]) {
			words = append(words[:i], words[i+1:]...)
			continue
		}

		i++
	}

	return words
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	nums := [26]int{}

	for i := range a {
		nums[a[i]-'a']++
		nums[b[i]-'a']--
	}

	for _, val := range nums {
		if val != 0 {
			return false
		}
	}

	return true
}
