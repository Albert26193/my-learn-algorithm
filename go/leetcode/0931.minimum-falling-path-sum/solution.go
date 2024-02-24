// Created by Albert at 2023/11/17 14:47
// leetgo: 1.3.8
// https://leetcode.cn/problems/minimum-falling-path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// f[i][j] ==> to matrix[i][j], minCost is f[i][j]
/*
	f[i][j] = minn(f[i-1][j], f[i-1][j-1], f[i-1][j+1]) + matrix[i][j]
*/

func minn(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

func minFallingPathSum(matrix [][]int) (ans int) {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	if col == 0 {
		return 0
	}

	var check = func(i int, j int) bool {
		if !(0 <= i && i <= row-1) {
			return false
		}

		if !(0 <= j && j <= col-1) {
			return false
		}

		return true
	}

	f := make([][]int, row)
	for i := 0; i < row; i++ {
		f[i] = make([]int, col)
	}

	for i := 0; i < col; i++ {
		f[0][i] = matrix[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			candidate := make([]int, 0)
			candidate = append(candidate, f[i-1][j])
			if check(i, j-1) {
				candidate = append(candidate, f[i-1][j-1])
			}

			if check(i, j+1) {
				candidate = append(candidate, f[i-1][j+1])
			}

			f[i][j] = minn(candidate...) + matrix[i][j]
		}
	}

	ans = minn(f[row-1]...)

	return ans
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	matrix := Deserialize[[][]int](ReadLine(reader))
	ans := minFallingPathSum(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
