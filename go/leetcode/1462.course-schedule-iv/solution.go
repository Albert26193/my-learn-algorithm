// Created by Albert at 2023/09/12 12:55
// leetgo: 1.3.7
// https://leetcode.cn/problems/course-schedule-iv/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) (ans []bool) {
	graph := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		fromNode := prerequisite[0]
		toNode := prerequisite[1]
		graph[fromNode] = append(graph[fromNode], toNode)
	}

	// fmt.Println(graph)

	var dfs func(int, []bool)
	dfs = func(x int, vis []bool) {
		vis[x] = true

		for _, toNode := range graph[x] {
			if !vis[toNode] {
				vis[toNode] = true
				dfs(toNode, vis)
			}
		}
	}

	for _, query := range queries {
		vis := make([]bool, numCourses)
		dfs(query[0], vis)
		if vis[query[1]] {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}

	}

	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// numCourses := Deserialize[int](ReadLine(stdin))
	// prerequisites := Deserialize[[][]int](ReadLine(stdin))
	// queries := Deserialize[[][]int](ReadLine(stdin))
	numCourses := 3
	prerequisites := [][]int{{1, 2}, {1, 0}, {2, 0}}
	queries := [][]int{{1, 0}, {1, 2}}
	ans := checkIfPrerequisite(numCourses, prerequisites, queries)

	fmt.Println("\noutput:", Serialize(ans))
}
