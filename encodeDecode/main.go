package main

import "fmt"

func main() {
	s := Solution{}

	strs := []string{"neet", "code", "love", "you"}
	encode := s.Encode(strs)
	fmt.Println("encode: ", encode)
	decode := s.Decode(encode)
	fmt.Println("decode: ", decode)
}

type Solution struct {
	delimiter []int
}

func (s *Solution) Encode(strs []string) string {
	delimiter := make([]int, len(strs))
	resp := ""

	for i, str := range strs {
		resp += str
		delimiter[i] = len(str)
	}

	s.delimiter = delimiter

	return resp
}

func (s *Solution) Decode(str string) []string {
	resp := make([]string, len(s.delimiter))

	carry := 0
	for i, val := range s.delimiter {
		resp[i] = str[carry:carry+val]
		carry += val
	}

	return resp
}
