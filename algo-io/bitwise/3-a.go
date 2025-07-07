package bitwise

// 10, 1 = 1010(2) => 1000 -> 8
//
//	&1101
func ClearPosition(n, pos int) int {
	bits := 0
	x := n
	for x >= 1 {
		x >>= 1
		bits++
	}

	return n & (1<<(bits+1) - 1<<pos - 1)
}
