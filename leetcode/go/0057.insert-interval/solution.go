// Created by Albert at 2023/09/12 17:07
// leetgo: 1.3.7
// https://leetcode.cn/problems/insert-interval/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

/*
for all intervals in origin intervals list, set current interval as C, newInterval as N:
 1. C's rightPoint < N's leftPoint: no cover ==> do nothing to C, push it to ans
 2. C's leftPoint > N's rightPoint: no cover ===> do nothing to C, push it to ans
 3. C's leftPoint <= N's rightPoint | C'rightPoint >= N's leftPoint ==> merge it

how to merge:

	set merged interval as M, M = [min(N's leftPoint, C's leftPoint), max(N's rightPoint, C's rightPoint)]
*/
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left := newInterval[0]
	right := newInterval[1]

	merged := false

	for _, interval := range intervals {
		// form low to high
		if right < interval[0] {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if left > interval[1] {
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	if !merged {
		ans = append(ans, []int{left, right})
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	newInterval := Deserialize[[]int](ReadLine(stdin))
	ans := insert(intervals, newInterval)

	fmt.Println("\noutput:", Serialize(ans))
}
