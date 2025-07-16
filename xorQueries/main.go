package main

func main() {

}

func xorQueries(arr []int, queries [][]int) []int {
	sums := make([]int, len(arr))
	sums[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		sums[i] = sums[i-1] ^ arr[i]
	}

	resp := make([]int, len(queries))

	for i, pairs := range queries {
		switch pairs[0] {
		case pairs[1]:
			resp[i] = arr[pairs[0]]
		case 0:
			resp[i] = sums[pairs[1]]
		default:
			resp[i] = sums[pairs[1]] ^ sums[pairs[0]-1]
		}
	}

	return resp
}
