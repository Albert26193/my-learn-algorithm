// Created by Albert's server at 2024/01/22 14:08
// leetgo: dev
// https://leetcode.cn/problems/spiral-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	row, col := len(matrix), len(matrix[0])

	var getNumber func(left, top, right, bottom int) []int
	getNumber = func(left, top, right, bottom int) []int {
		if left > right || top > bottom {
			// fmt.Println("here")
			return []int{}
		}

		// get top
		res := make([]int, 0)
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		// get right
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		// get bottom
		for i := right - 1; i >= left && top < bottom; i-- {
			res = append(res, matrix[bottom][i])
		}

		//get left
		for i := bottom - 1; i >= top+1 && left < right; i-- {
			res = append(res, matrix[i][left])
		}

		res = append(res, getNumber(left+1, top+1, right-1, bottom-1)...)
		// fmt.Println(res)
		return res
	}

	return getNumber(0, 0, col-1, row-1)
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
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
