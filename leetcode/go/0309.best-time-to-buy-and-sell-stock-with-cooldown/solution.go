// Created by Albert at 2023/09/30 20:22
// leetgo: 1.3.8
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/

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
	/*
	   dp[i][j]: i ==> prices[0]...prices[i - 1]
	           j ==> 0 or 1 or 2
	       i == 0: have a stock to sell
	       i == 1: not have a stock, but is cooldown
	       i == 2: not have a stock, and not in cooldown
	*/

	/*
	   dp[i + 1][0] = dp[i][0]
	   dp[i + 1][0] = dp[i][2] - prices[i]

	   dp[i + 1][1] = dp[i][0] + prices[i]

	   dp[i + 1][2] = dp[i][2]
	   dp[i + 1][2] = dp[i][1]
	*/

	days := len(prices)
	dp := make([][3]int, days+1)

	dp[0][0] = -1e5
	dp[0][1] = -1e5
	dp[0][2] = 0
	for i := 0; i < days; i++ {
		dp[i+1][0] = maxx(dp[i][0], dp[i][2]-prices[i])
		dp[i+1][1] = dp[i][0] + prices[i]
		dp[i+1][2] = maxx(dp[i][1], dp[i][2])
	}

	return maxx(dp[days][1], dp[days][2])
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
