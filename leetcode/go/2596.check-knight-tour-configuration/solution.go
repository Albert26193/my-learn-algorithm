// Created by Albert at 2023/09/13 08:46
// leetgo: 1.3.7
// https://leetcode.cn/problems/check-knight-tour-configuration/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkValidGrid(grid [][]int) bool {
	col := len(grid)
	if col == 0 {
		return false
	}

	row := len(grid[0])
	if row == 0 {
		return false
	}

	if grid[0][0] != 0 {
		return false
	}

	checkRange := func(x, y int) bool {
		if (x < 0) || (x >= row) {
			return false
		}

		if (y < 0) || (y >= col) {
			return false
		}

		return true
	}

	directions := [][]int{
		{-2, 1},
		{-1, 2},
		{1, 2},
		{2, 1},
		{2, -1},
		{1, -2},
		{-1, -2},
		{-2, -1},
	}

	startPosition := []int{0, 0}
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == 0 {
				startPosition = []int{x, y}
			}
			break
		}
	}

	vis := make([]int, row*col)
	var dfs func([]int, int)
	dfs = func(currentPosition []int, currentVal int) {
		for _, direction := range directions {
			vis[currentVal] = 1
			nextPostion := []int{
				currentPosition[0] + direction[0],
				currentPosition[1] + direction[1],
			}

			if !checkRange(nextPostion[0], nextPostion[1]) {
				continue
			}

			if grid[nextPostion[0]][nextPostion[1]] != currentVal+1 {
				continue
			}

			nextVal := currentVal + 1
			if vis[nextVal] == 1 {
				continue
			}

			if nextVal < row*col {
				dfs(nextPostion, nextVal)
			}
		}
	}

	dfs(startPosition, 0)

	fmt.Println(vis)
	if vis[row*col-1] == 1 {
		return true
	}

	return false
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// grid := Deserialize[[][]int](ReadLine(stdin))
	// [[24,11,22,17,4],[21,16,5,12,9],[6,23,10,3,18],[15,20,1,8,13],[0,7,14,19,2]]
	grid := [][]int{
		{24, 11, 22, 17, 4},
		{21, 16, 5, 12, 9},
		{6, 23, 10, 3, 18},
		{15, 20, 1, 8, 13},
		{0, 7, 14, 19, 2},
	}

	ans := checkValidGrid(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
