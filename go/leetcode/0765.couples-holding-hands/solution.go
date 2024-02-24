// Created by Albert's server at 2024/01/21 00:11
// leetgo: dev
// https://leetcode.cn/problems/couples-holding-hands/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSwapsCouples(row []int) (ans int) {
	fa := make(map[int]int)

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}

		return fa[x]
	}

	var merge = func(x int, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			fa[fx] = fy
		}
	}

	n := len(row)
	m := n / 2

	for i := 0; i < m; i++ {
		fa[i] = i
	}

	for i := 0; i < n; i += 2 {
		merge(row[i]/2, row[i+1]/2)
	}

	cnt := 0
	for i := 0; i < m; i++ {
		if i == find(i) {
			cnt += 1
		}
	}

	return m - cnt
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
	row := Deserialize[[]int](ReadLine(in))
	ans := minSwapsCouples(row)

	fmt.Println("\noutput:", Serialize(ans))
}
