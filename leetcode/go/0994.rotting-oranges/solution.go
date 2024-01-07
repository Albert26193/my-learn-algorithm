// Created by Albert's server at 2024/01/07 14:30
// leetgo: dev
// https://leetcode.cn/problems/rotting-oranges/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func orangesRotting(grid [][]int) (ans int) {
	row := len(grid)
	col := len(grid[0])

	var d = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var check = func(x int, y int) bool {
		if x < 0 || x >= row {
			return false
		}

		if y < 0 || y >= col {
			return false
		}
		return true
	}

	vis := make([]int, row*col)
	q := make([]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				index := i*col + j
				q = append(q, index)
				vis[index] = 1
			}
		}
	}

	step := 0

	isFirstLevel := true
	for len(q) > 0 {
		levelSize := len(q)

		if isFirstLevel {
			step -= 1
			isFirstLevel = false
		}

		for i := 0; i < levelSize; i++ {
			head := q[0]
			q = q[1:]

			x := head / col
			y := head % col

			for j := 0; j < 4; j++ {
				nx := x + d[j][0]
				ny := y + d[j][1]

				if !check(nx, ny) {
					continue
				}

				if grid[nx][ny] != 1 {
					continue
				}

				next := nx*col + ny

				if vis[next] == 1 {
					continue
				}

				q = append(q, next)
				vis[next] = 1
			}
		}

		step += 1
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if vis[i*col+j] != 1 && grid[i][j] != 0 {
				return -1
			}
		}
	}

	return step
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
	grid := Deserialize[[][]int](ReadLine(in))
	ans := orangesRotting(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
