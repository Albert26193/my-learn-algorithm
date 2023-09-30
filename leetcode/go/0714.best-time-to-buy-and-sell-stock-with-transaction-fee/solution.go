// Created by Albert at 2023/09/30 22:31
// leetgo: 1.3.8
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

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

func maxProfit(prices []int, fee int) (ans int) {
	days := len(prices)
	dp := make([][2]int, days+1)

	// dp[i][0] ==> have a stock
	// dp[i][1] ==> not have a stock

    dp[0][0] = -1e5
    dp[0][1] = 0
	for i := 0; i < days; i++ {
		dp[i+1][0] = maxx(dp[i][1]-prices[i], dp[i][0])
		dp[i+1][1] = maxx(dp[i][0]+prices[i]-fee, dp[i][1])
	}

	return dp[days][1]
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
	fee := Deserialize[int](ReadLine(reader))
	ans := maxProfit(prices, fee)

	fmt.Println("\noutput:", Serialize(ans))
}
