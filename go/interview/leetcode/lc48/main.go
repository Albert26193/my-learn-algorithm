package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(ma [][]int) {
	if len(ma) == 0 || len(ma[0]) == 0 {
		return
	}

	row, col := len(ma), len(ma[0])

	for i := 0; i < row/2; i++ {
		for j := 0; j < (col+1)/2; j++ {
			temp := ma[i][j]
			ma[i][j] = ma[row-1-j][i]
			ma[row-1-j][i] = ma[row-1-i][row-1-j]
			ma[row-1-i][row-1-i] = ma[j][row-1-i]
			ma[j][row-1-i] = temp
		}
	}
}
