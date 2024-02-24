// Created by Albert at 2023/04/13 11:50
// https://leetcode.cn/problems/matchsticks-to-square/

package main

import (
	"fmt"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
// backTrack index means: from index To the end of matchsticks array, 
// all of the matchsticks can be put into eligible edge
func backTrack(matchsticks []int, stickIndex int, edges []int, target int) bool {
	if stickIndex >= len(matchsticks) {
		return true
	}

	for i := 0; i < 4; i++ {
		if edges[i]+matchsticks[stickIndex] > target {
			continue
		}
		edges[i] += matchsticks[stickIndex]
		if backTrack(matchsticks, stickIndex+1, edges, target) {
			return true
		}
		edges[i] -= matchsticks[stickIndex]
	}

	return false
}

func makesquare(matchsticks []int) bool {
	sticksSum := 0
	for i := range matchsticks {
		sticksSum += matchsticks[i]
	}

	if sticksSum%4 != 0 {
		return false
	}

	target := sticksSum / 4
	// sort.Ints(matchsticks)

	// why descending sort is better than ascending sort in TIME ?
	// pruning in advance
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	edges := make([]int, 4)
	return backTrack(matchsticks, 0, edges, target)
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// matchsticks := Deserialize[[]int](ReadLine(stdin))
	matchsticks := []int{1, 1, 2, 2, 2}
	matchsticks = []int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}
	ans := makesquare(matchsticks)
	fmt.Println("output: " + Serialize(ans))
}
