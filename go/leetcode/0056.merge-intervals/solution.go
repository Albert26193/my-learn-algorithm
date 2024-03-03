// Created by Albert at 2023/09/12 16:23
// leetgo: 1.3.7
// https://leetcode.cn/problems/merge-intervals/

package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"sort"

	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
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
*/

// TODO: implement it again
func merge(intervals [][]int) (ans [][]int) {
	if len(intervals) == 0 {
		return ans
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for _, inter := range intervals {
		if len(merged) == 0 {
			merged = append(merged, inter)
			continue
		}

		// right ==> last merged one's right boundary
		lastMerged := merged[len(merged)-1]
		rightPivot := lastMerged[1]

		cur := inter
		// current left < rightPivot
		if inter[0] <= rightPivot {
			cur = []int{lastMerged[0], maxx(inter[1], lastMerged[1])}
			merged = merged[:len(merged)-1]
		}

		merged = append(merged, cur)
	}

	// fmt.Println(intervals)
	return merged
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
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

	intervals := Deserialize[[][]int](ReadLine(in))

	ans := merge(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
