// Created by Albert's server at 2024/01/13 23:22
// leetgo: dev
// https://leetcode.cn/problems/n-queens-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func totalNQueens(n int) (ans int) {
	col := make([]int, n)
	dg := make([]int, 2*n)
	rdg := make([]int, 2*n)

	var dfs func(r int)

	dfs = func(r int) {
		if r == n-1 {
			ans += 1
			return
		}

		for c := 0; c < n; c++ {
			if col[c] == 1 || dg[c+r] == 1 || rdg[c-r+n] == 1 {
				continue
			}

			col[c], dg[c+r], rdg[c-r+n] = 1, 1, 1
			dfs(r + 1)
			col[c], dg[c+r], rdg[c-r+n] = 0, 0, 0
		}
	}

	dfs(0)
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
	n := Deserialize[int](ReadLine(in))
	ans := totalNQueens(n)

	fmt.Println("\noutput:", Serialize(ans))
}
