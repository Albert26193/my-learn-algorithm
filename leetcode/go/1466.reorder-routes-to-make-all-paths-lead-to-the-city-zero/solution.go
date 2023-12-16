// Created by Albert's server at 2023/12/07 12:36
// leetgo: dev
// https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// TODO: check official solutions
func minReorder(n int, connections [][]int) (ans int) {
	dia := make([][]int, n)
	recordDirection := make(map[string]bool)

	for _, v := range connections {
		start := v[0]
		end := v[1]
		dia[start] = append(dia[start], end)
		dia[end] = append(dia[end], start)
		direction := strconv.Itoa(start) + "#" + strconv.Itoa(end)
		recordDirection[direction] = true
	}

	vis := make([]int, n)
	var dfs func(int, int)
	dfs = func(curNode int, faNode int) {
		nextNodes := dia[curNode]
		if vis[curNode] == 1 {
			return
		}

		vis[curNode] = 1

		relation := strconv.Itoa(curNode) + "#" + strconv.Itoa(faNode)
		if !recordDirection[relation] {
			// fmt.Println(curNode, faNode, "fff")
			ans += 1
		}

		for _, node := range nextNodes {
			dfs(node, curNode)
		}
	}

	dfs(0, -1)
	return ans - 1
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
	connections := Deserialize[[][]int](ReadLine(in))
	// fmt.Println(n, connections)
	ans := minReorder(n, connections)

	fmt.Println("\noutput:", Serialize(ans))
}
