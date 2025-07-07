package main

import "fmt"

func main() {
	n := 4
	trust := [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {1, 4}, {4, 3}, {4, 1}}
	fmt.Println(findJudge(n, trust)) // Output: -1
}

func findJudge(n int, trust [][]int) int {
	judge := - 1
	in := make([]int, n + 1)
	out	:= make([]int, n + 1)

	for _, val := range trust {
		in[val[1]]++
		out[val[0]]++
	}

	for i := 1; i <= n; i++ {
		if in[i] == n - 1 && out[i] == 0 {
			judge = i
		}
	}

	return judge
}
