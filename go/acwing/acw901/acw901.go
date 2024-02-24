package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./testcase.txt")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	in := bufio.NewReader(file)

	defer file.Close()

	var r, c int
	fmt.Fscan(in, &r, &c)

	grid := make([][]int, r)
	f := make([][]int, r)

	for i := 0; i < r; i++ {
		grid[i] = make([]int, c)
		f[i] = make([]int, c)
		for j := 0; j < c; j++ {
			var x int
			fmt.Fscan(in, &x)
			grid[i][j] = x
		}
	}

	// fmt.Println(grid, r, c)
	// return

	var check = func(x int, y int) bool {
		if x < 0 || x >= r || y < 0 || y >= c {
			return false
		}
		return true
	}

	var d = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	ans := 0
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		if !check(x, y) || f[x][y] != 0 {
			return
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+d[i][0], y+d[i][1]
			if !check(nx, ny) || grid[nx][ny] >= grid[x][y] {
				continue
			}

			// first dfs
			dfs(nx, ny)

			// then update
			f[x][y] = maxx(f[x][y], f[nx][ny]+1)
			ans = maxx(ans, f[x][y])
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			dfs(i, j)
		}
	}

	fmt.Println(ans + 1)
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
