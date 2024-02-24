// Created by Albert's server at 2024/01/22 13:48
// leetgo: dev
// https://leetcode.cn/problems/set-matrix-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func setZeroes(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])

	recRow := false
	recCol := false

	// record first row
	for i := 0; i < col; i++ {
		if matrix[0][i] == 0 {
			recRow = true
		}
	}

	// record first col
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			recCol = true
		}
	}

	// put flag into first row and first col
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	// put into new matrix
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// if have zero in first row or first col
	if recRow {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}

	if recCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	matrix := Deserialize[[][]int](ReadLine(in))
	setZeroes(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
