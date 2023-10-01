// Created by Albert at 2023/10/01 11:21
// leetgo: 1.3.8
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int) (ans int) {
	days := len(prices)
	rightMax := make([]int, days)

	maxNum := -5
	maxDiff := 0
	for i := days - 1; i >= 0; i-- {
		maxNum = maxx(maxNum, prices[i])
		rightMax[i] = maxNum
		maxDiff = maxx(maxDiff, maxNum-prices[i])
	}

	// fmt.Println(rightMax)
	return maxDiff
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

	prices := Deserialize[[]int](ReadLine(reader))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
