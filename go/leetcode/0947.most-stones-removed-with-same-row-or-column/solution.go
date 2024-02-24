// Created by Albert's server at 2024/01/16 21:51
// leetgo: dev
// https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func removeStones(stones [][]int) int {
	var fa = make(map[int]int)

	var find func(int) int
	find = func(x int) int {
		if _, ok := fa[x]; !ok {
			fa[x] = x
		}

		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	var merge = func(x int, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		fa[fx] = fy
	}

	var getCnt = func() int {
		rec := 0
		for k := range fa {
			if fa[k] == k {
				rec += 1
			}
		}
		return rec
	}

	const offset = 100001

	n := len(stones)

	for _, v := range stones {
		merge(v[0]+offset, v[1])
	}

	rec := getCnt()
	return n - rec
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
	stones := Deserialize[[][]int](ReadLine(in))
	ans := removeStones(stones)

	fmt.Println("\noutput:", Serialize(ans))
}
