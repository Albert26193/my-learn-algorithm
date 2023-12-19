// Created by Albert at 2023/11/23 16:22
// leetgo: 1.3.8
// https://leetcode.cn/problems/word-search/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])

	var d = [4][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	check := func(x int, y int) bool {
		if x < 0 || x >= row {
			return false
		}
		if y < 0 || y >= col {
			return false
		}
		return true
	}

	var dfs func(int, int, int)
	n := len(word)

	vis := make([]bool, row*col)
	hasWord := false

	dfs = func(depth int, x int, y int) {
		if depth >= n-1 {
			hasWord = true
			return
		}

		if !check(x, y) || vis[x*col+y] || board[x][y] != word[depth] {
			return
		}

		vis[x*col+y] = true
		for i := 0; i < 4; i++ {
			nx := x + d[i][0]
			ny := y + d[i][1]
			if !check(nx, ny) {
				continue
			}

			if vis[nx*col+ny] {
				continue
			}

			if board[nx][ny] != word[depth+1] {
				continue
			}

			dfs(depth+1, nx, ny)
		}

		vis[x*col+y] = false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				hasWord = false
				vis = make([]bool, row*col)
				dfs(0, i, j)

				if hasWord {
					return true
				}
			}
		}
	}

	return false
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
	board := Deserialize[[][]byte](ReadLine(in))
	word := Deserialize[string](ReadLine(in))
	ans := exist(board, word)

	fmt.Println("\noutput:", Serialize(ans))
}
