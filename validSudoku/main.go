package main

import "fmt"

func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	// fmt.Println(isValidSudoku(board))

	// board2 := [][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	// fmt.Println(isValidSudoku(board2))

	board3 := [][]byte{
		{'.', '.', '5', '.', '.', '.', '.', '.', '.'},
		{'1', '.', '.', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '6', '.', '.', '3', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '1', '5', '2', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
		{'.', '.', '6', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(board3))

}

func isValidSudoku(board [][]byte) bool {
	row := make(map[int]bool)
	col := make(map[int]bool)

	for i := range 9 {
		for j := range 9 {
			val := board[i][j]

			if val == '.' {
				continue
			}

			board[i][j] = '.'

			// if !row[i] {
			fmt.Println("row", i, j, string(val))
			if checkRow(board, i, val) {
				return false
			}
			// }

			// if !col[j] {
			fmt.Println("col", i, j, string(val))
			if checkCol(board, j, val) {
				return false
			}
			// }

			row[i] = true
			col[j] = true
		}
	}

	return true
}

func checkRow(board [][]byte, row int, val byte) bool {
	for i := range 9 {
		// fmt.Println(string(board[row][i]), string(val))
		if board[row][i] == val {
			return true
		}
	}

	return false
}

func checkCol(board [][]byte, col int, val byte) bool {
	for i := range 9 {
		fmt.Println("check col", string(board[i][col]), string(val))
		if board[i][col] == val {
			return true
		}
	}

	return false
}
