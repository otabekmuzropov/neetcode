package main

func main() {

}

func isValid(s string) bool {
	// stc := NewStack(len(s))
	// for i := range s {
	// 	if open[s[i]] {

	// 	}
	// }

	return false
}

type Stack struct {
	stack []byte
}

func NewStack(length int) Stack {
	return Stack{
		stack: make([]byte, length),
	}
}

func (s *Stack) Push(val byte) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() {
	i := len(s.stack) - 1
	s.stack = s.stack[:i]
}

var open map[byte]bool = map[byte]bool{
	'{': true,
	'(': true,
	'[': true,
	'}': false,
	')': false,
	']': false,
}
