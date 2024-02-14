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
type position struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {
	var getLevel func(start position, end position) []int
	getLevel = func(start position, end position) []int {
		if start.x > end.x || start.y > end.y {
			fmt.Println(start.x, end.x)
			return []int{}
		}

		if start.x == end.x && start.y == end.y {
			return []int{matrix[start.x][start.y]}
		}

		top, left, bottom, right := start.x, start.y, end.x, end.y
		cur := make([]int, 0)

		// top
		for i := left; i <= right; i++ {
			cur = append(cur, matrix[top][i])
		}

		// right
		for i := top + 1; i <= bottom; i++ {
			cur = append(cur, matrix[i][right])
		}

		// bottom
		for i := right - 1; i >= left && top < bottom; i-- {
			cur = append(cur, matrix[bottom][i])
		}

		// left
		for i := bottom - 1; i > top && left < right; i-- {
			cur = append(cur, matrix[i][left])
		}

		child := getLevel(position{top + 1, left + 1}, position{bottom - 1, right - 1})
		cur = append(cur, child...)
		return cur
	}

	row, col := len(matrix), len(matrix[0])
	p1, p2 := position{0, 0}, position{row - 1, col - 1}
	ans := getLevel(p1, p2)
	// fmt.Println(ans)
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
	matrix := Deserialize[[][]int](ReadLine(in))
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
