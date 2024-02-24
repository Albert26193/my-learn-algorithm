// Created by Albert's server at 2024/02/16 20:16
// leetgo: dev
// https://leetcode.cn/problems/maximal-rectangle/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximalRectangle(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])

	mat := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		mat[i] = make([]int, col+1)
	}

	ans := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if matrix[i-1][j-1] == '0' {
				mat[i][j] = 0
			} else {
				mat[i][j] = mat[i-1][j] + 1
			}
		}

		ans = maxx(ans, largestRec(mat[i]))
		// fmt.Println(largestRec(mat[i]), i, "ffffffffff")
	}

	// fmt.Println(mat)

	return ans
}

func maxx(nums ...int) int {
	res := nums[0]

	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
}

func largestRec(h []int) int {
	ans := 0
	ls, rs := make([]int, 0), make([]int, 0)
	lm, rm := make(map[int]int), make(map[int]int)
	n := len(h)

	for i := 0; i < n; i++ {
		for len(ls) > 0 && h[ls[len(ls)-1]] >= h[i] {
			ls = ls[:len(ls)-1]
		}

		if len(ls) == 0 {
			lm[i] = -1
		} else {
			lm[i] = ls[len(ls)-1]
		}

		ls = append(ls, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(rs) > 0 && h[rs[len(rs)-1]] >= h[i] {
			rs = rs[:len(rs)-1]
		}

		if len(rs) == 0 {
			rm[i] = n
		} else {
			rm[i] = rs[len(rs)-1]
		}

		rs = append(rs, i)
	}

	// fmt.Println(lm, rm)

	for i := 0; i < n; i++ {
		square := (rm[i] - lm[i] - 1) * h[i]
		ans = maxx(ans, square)
	}

	return ans
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
	matrix := Deserialize[[][]byte](ReadLine(in))
	ans := maximalRectangle(matrix)

	// fmt.Println(largestRec([]int{2, 1, 5, 6, 2, 3}))
	// fmt.Println(largestRec([]int{2, 4}))
	// fmt.Println(largestRec([]int{3, 1, 3, 2, 2}))
	fmt.Println("\noutput:", Serialize(ans))
}
