// Created by Albert at 2023/10/02 11:10
// leetgo: 1.3.8
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 1. redo it; 2. 2 state is enough, not need 3 state
func maxx(a ...int) int {
	res := a[0]
	for _, num := range a {
		if num > res {
			res = num
		}
	}

	return res
}

func maxProfit(prices []int) (ans int) {
	/*
	   dp[i][0] ==> prices[0]...prices[i-1]
	   dp[i][0] ==> 0: not buy a stock
	   dp[i][1] ==> 1: buy a stock, but not sell it
	   dp[i][2] ==> 2: buy a stock, and sell it
	*/

	days := len(prices)
	dp := make([][3]int, days+1)

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	for i := range prices {
		dp[i+1][0] = dp[i][0]
		dp[i+1][1] = maxx(dp[i][1], dp[i][0]-prices[i], dp[i][2]-prices[i])
		dp[i+1][2] = maxx(dp[i][2], dp[i][1]+prices[i])
	}

	return maxx(dp[days][0], dp[days][1], dp[days][2])
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
