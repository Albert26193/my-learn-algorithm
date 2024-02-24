// Created by Albert at 2023/09/21 11:42
// leetgo: 1.3.8
// https://leetcode.cn/problems/minimum-height-trees/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findMinHeightTrees(n int, edges [][]int) (ans []int) {
	// BFS 剥洋葱,一层一层拿掉
	// 1. 建立邻接表, 计算某个节点的 "度"
	// 2. 利用BFS, 度为1就加入队列; 同时,建立一个哈希表,记录 level: [node1, node2 ...]
	// 3. 读出最后一层拿走的叶子节点,就是所需的节点

	// 1. 建立邻接表, 计算某个节点的 "度"
	if n == 1 {
		return []int{0}
	}

	degree := make([]int, n)
	adjency := make([][]int, n)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		degree[from] += 1
		degree[to] += 1

		adjency[from] = append(adjency[from], to)
		adjency[to] = append(adjency[to], from)
	}

	// 2. 利用BFS, 度为1就加入队列; 同时,建立一个哈希表,记录 level: [node1, node2 ...]
	levelMap := make(map[int][]int)
	queue := make([]int, 0)

	// degree = 1 ==> push it to queue
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	levelNum := 0
	for len(queue) > 0 {
		queueLength := len(queue)
		levelNum += 1
		for i := 0; i < queueLength; i++ {
			curNode := queue[0]
			queue = queue[1:]
			levelMap[levelNum] = append(levelMap[levelNum], curNode)

			for _, nextNode := range adjency[curNode] {
				degree[nextNode] -= 1
                // 为什么这里不会有重复加入队列的问题?
                // 因为之前加入进队列的, 入度会变成0
				if degree[nextNode] == 1 {
					queue = append(queue, nextNode)
				}
			}
		}
	}

	return levelMap[levelNum]
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	n := Deserialize[int](ReadLine(reader))
	edges := Deserialize[[][]int](ReadLine(reader))
	ans := findMinHeightTrees(n, edges)

	fmt.Println("\noutput:", Serialize(ans))
}
