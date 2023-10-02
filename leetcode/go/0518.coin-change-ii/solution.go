// Created by Albert at 2023/10/02 16:52
// leetgo: 1.3.8
// https://leetcode.cn/problems/coin-change-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
func change(amount int, coins []int) (ans int) {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
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

	amount := Deserialize[int](ReadLine(reader))
	coins := Deserialize[[]int](ReadLine(reader))
	ans := change(amount, coins)

	fmt.Println("\noutput:", Serialize(ans))
}
