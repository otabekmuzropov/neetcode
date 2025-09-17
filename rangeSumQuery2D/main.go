package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{1, 2, 3, 4},
	}
	nm := Constructor(matrix)
	fmt.Println(nm.grid)
}

type NumMatrix struct {
	grid [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	column := len(matrix[0])
	grid := make([][]int, row+1)

	for i := range row + 1 {
		grid[i] = make([]int, column+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			fmt.Println(i, j)
			grid[i][j] = matrix[i-1][j-1] + grid[i-1][j] + grid[i][j-1] - grid[i-1][j-1]
		}
	}

	return NumMatrix{
		grid: grid,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var conj1, conj2, adj int
	if row1 > 0 {
		conj1 = this.grid[row1-1][col2]
	}

	if col1 > 0 {
		conj2 = this.grid[row2][col1-1]
	}

	if row1 > 0 && col1 > 0 {
		adj = this.grid[row1-1][col1-1]
	}

	return this.grid[row2][col2] - conj1 - conj2 + adj
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
