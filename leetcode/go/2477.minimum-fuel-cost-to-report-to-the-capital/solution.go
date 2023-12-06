// Created by Albert's server at 2023/12/05 12:23
// leetgo: dev
// https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// calculate all cost as c
func minimumFuelCost(roads [][]int, seats int) (ans int64) {
	fa := make(map[int]int)
	roadsCount := len(roads) + 1

	dis := make([]int, roadsCount)

	for i := 0; i < roadsCount; i++ {
		dis[i] = -1
	}

	roadsRecord := make([][]int, roadsCount)

	for _, road := range roads {
		start := road[0]
		end := road[1]
		roadsRecord[start] = append(roadsRecord[start], end)
		roadsRecord[end] = append(roadsRecord[end], start)
	}

	// BFS
	queue := make([]int, 0)
	queue = append(queue, 0)
	level := 0
	maxLevel := 0
	disToCity := make([][]int, roadsCount)

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			head := queue[0]
			queue = queue[1:]
			dis[head] = level
			disToCity[level] = append(disToCity[level], head)
			maxLevel = level

			for _, nextCity := range roadsRecord[head] {
				if dis[nextCity] != -1 {
					continue
				}
				fa[nextCity] = head
				queue = append(queue, nextCity)
			}
		}
		level += 1
	}

	// dis array is key for calculate
	curNum := make([]int, roadsCount)
	for i := 1; i < roadsCount; i++ {
		curNum[i] = 1
	}

	carsCount := 0
	for i := maxLevel; i > 0; i-- {
		for _, city := range disToCity[i] {
			carsCount += curNum[city] / seats
			peopleRest := curNum[city] % seats
			if peopleRest > 0 {
				carsCount += 1
			}

			curNum[fa[city]] += curNum[city]
		}
	}

	return int64(carsCount)
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

	roads := Deserialize[[][]int](ReadLine(in))
	seats := Deserialize[int](ReadLine(in))
	ans := minimumFuelCost(roads, seats)

	fmt.Println("\noutput:", Serialize(ans))
}
