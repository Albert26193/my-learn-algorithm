// Created by Albert at 2023/09/12 16:23
// leetgo: 1.3.7
// https://leetcode.cn/problems/merge-intervals/

package main

import (
	// "bufio"
	"fmt"
	// "os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	startPoint := intervals[0][0]
	endPoint := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= endPoint {
			if intervals[i][1] > endPoint {
				endPoint = intervals[i][1]
			}
		} else {
			ans = append(ans, []int{startPoint, endPoint})
			startPoint = intervals[i][0]
			endPoint = intervals[i][1]
		}
	}

	ans = append(ans, []int{startPoint, endPoint})

	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// intervals := Deserialize[[][]int](ReadLine(stdin))

	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	ans := merge(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
