package main

import "fmt"

func main() {

}

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		for _, ii := range row {
			fmt.Println(ii)
		}
	}

	return 1 == 1
}
