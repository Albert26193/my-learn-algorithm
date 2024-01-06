// Created by Albert's server at 2024/01/06 20:47
// leetgo: dev
// https://leetcode.cn/problems/graph-valid-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func validTree(n int, edges [][]int) bool {

	edgeCount := len(edges)
	if edgeCount != (n - 1) {
		return false
	}

	vis := make([]int, n)

	g := make([][]int, n)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)

		if from == to {
			return false
		}
	}

	var dfs func(cur int)
	dfs = func(cur int) {
		nextNodes := g[cur]

		vis[cur] = 1
		for _, node := range nextNodes {
			if vis[node] != 1 {
				dfs(node)
			}
		}
	}

	dfs(0)

	for i := 0; i < n; i++ {
		if vis[i] != 1 {
			return false
		}
	}

	return true
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
	edges := Deserialize[[][]int](ReadLine(in))
	ans := validTree(n, edges)

	fmt.Println("\noutput:", Serialize(ans))
}
