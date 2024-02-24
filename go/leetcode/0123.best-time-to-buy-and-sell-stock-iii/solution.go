// Created by Albert at 2023/09/30 23:33
// leetgo: 1.3.8
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/

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
	   dp[i][0] ==> prices[0]...prices[i - 1]

	   dp[i][0] ==> not buy 1 and not buy 2
	   dp[i][1] ==> buy 1 and not sell
	   dp[i][2] ==> buy 1 and sell it, but not buy 2
	   dp[i][3] ==> buy 1 and sell it, and buy 2, but not sell it
	   dp[i][4] ==> buy 1 and sell it, buy 2 and sell it

	   dp[i + 1][0] = dp[i][0]
	   dp[i + 1][1] = dp[i][0] - prices[i]
	   dp[i + 1][1] = dp[i][1]
	   dp[i + 1][2] = dp[i][1] + prices[i]
	   dp[i + 1][2] = dp[i][2]
	   dp[i + 1][3] = dp[i][2] - prices[i]
	   dp[i + 1][3] = dp[i][3]
	   dp[i + 1][4] = dp[i][3] + prices[i]
	   dp[i + 1][4] = dp[i][4]
	*/

	days := len(prices)
	dp := make([][5]int, days+1)

	dp[0][0] = 0
	dp[0][1] = -1e5
	dp[0][2] = -1e5
	dp[0][3] = -1e5
	dp[0][4] = -1e5
	for i := 0; i < days; i++ {
		dp[i+1][0] = dp[i][0]
		dp[i+1][1] = maxx(dp[i][0]-prices[i], dp[i][1])
		dp[i+1][2] = maxx(dp[i][1]+prices[i], dp[i][2])
		dp[i+1][3] = maxx(dp[i][2]-prices[i], dp[i][3])
		dp[i+1][4] = maxx(dp[i][3]+prices[i], dp[i][4])
	}

	return maxx(maxx(dp[days][4], dp[days][2]), dp[days][0])
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
