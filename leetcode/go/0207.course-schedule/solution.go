// Created by Albert's server at 2024/01/06 20:22
// leetgo: dev
// https://leetcode.cn/problems/course-schedule/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	ind := make([]int, numCourses)

	for _, pair := range prerequisites {
		pre := pair[1]
		next := pair[0]

		g[pre] = append(g[pre], next)
		// g[next] = append(g[next], pre)
		ind[next] += 1
	}

	q := make([]int, 0)
	vis := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if ind[i] == 0 {
			q = append(q, i)
			vis[i] = 1
		}
	}

	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			head := q[len(q)-1]
			q = q[:len(q)-1]

			nextNodes := g[head]
			for _, node := range nextNodes {
				ind[node] -= 1

				if ind[node] == 0 {
					q = append(q, node)
					vis[node] = 1
				}
			}
		}
	}

	fmt.Println(ind, vis)
	for i := 0; i < numCourses; i++ {
		if (ind[i] != 0) || (vis[i] != 1) {
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
	numCourses := Deserialize[int](ReadLine(in))
	prerequisites := Deserialize[[][]int](ReadLine(in))
	ans := canFinish(numCourses, prerequisites)

	fmt.Println("\noutput:", Serialize(ans))
}
