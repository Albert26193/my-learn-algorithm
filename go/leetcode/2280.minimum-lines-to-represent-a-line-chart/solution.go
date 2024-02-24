// Created by Albert at 2023/09/03 11:36
// leetgo: 1.3.7
// https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/

package main

import (
	// "bufio"
	"fmt"
	"sort"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumLines(stockPrices [][]int) (ans int) {
	dateCount := len(stockPrices)

	if dateCount <= 2 {
		return dateCount - 1
	}

	lineCount := 1

	sort.Slice(stockPrices, func(a, b int) bool {
		return stockPrices[a][0] < stockPrices[b][0]
	})

	for i := 2; i < dateCount; i++ {
		dx := stockPrices[i-1][0] - stockPrices[i-2][0]
		dy := stockPrices[i-1][1] - stockPrices[i-2][1]

		ndx := stockPrices[i][0] - stockPrices[i-1][0]
		ndy := stockPrices[i][1] - stockPrices[i-1][1]

		if dx*ndy != ndx*dy {
			lineCount += 1
		}
	}
	return lineCount
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// stockPrices := Deserialize[[][]int](ReadLine(stdin))
	stockPrices := [][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}
	stockPrices = [][]int{{1, 1}, {2, 2}, {3, 1}}
	ans := minimumLines(stockPrices)

	fmt.Println("\noutput:", Serialize(ans))
}
