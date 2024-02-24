// Created by Albert at 2023/11/07 13:19
// leetgo: 1.3.8
// https://leetcode.cn/problems/coin-change/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
	dp[i][j]:
	i: [0, ..., i-1]
	j: sum == j
	dp[i][j] == m ==> m is the ans for least count

	not choose current: dp[i][j] = dp[i - 1][j]
	choose current    : d[i][j] = dp[i - 1][j - coins[i - 1]] + 1
*/

func minn(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

// func coinChange(coins []int, amount int) (ans int) {
// 	const inf = int(0x3f3f3f3f)
// 	coinsCount := len(coins)
// 	dp := make([][]int, coinsCount+1)
// 	for i := 0; i <= coinsCount; i++ {
// 		dp[i] = make([]int, amount+1)
// 		for j := 0; j <= amount; j++ {
// 			dp[i][j] = inf
// 		}
// 	}
//
// 	for i := 0; i <= coinsCount; i++ {
// 		dp[i][0] = 0
// 	}
//
// 	for i := 1; i <= amount; i++ {
// 		dp[0][i] = inf
// 	}
//
// 	ans = inf
// 	for i := 1; i <= coinsCount; i++ {
// 		for j := 0; j <= amount; j++ {
// 			currentCoin := coins[i-1]
// 			for k := 0; k <= j/currentCoin; k++ {
// 				dp[i][j] = minn(dp[i][j], dp[i-1][j-k*currentCoin]+k)
// 			}
// 		}
//
// 		ans = minn(ans, dp[i][amount])
// 	}
//
// 	if ans == inf {
// 		return -1
// 	}
//
// 	return ans
// }

// 状态压缩
/*
	dp[i] ==> sum == i, min counts == dp[i]

	dp[i] = dp[i - coins[0]] + 1
	dp[i] = dp[i - coins[1]] + 1
	...
	dp[i] = dp[i - coins[len - 1] ] + 1

*/
func coinChange(coins []int, amount int) (ans int) {
	const inf = int(0x3f3f3f3f)
	coinsCount := len(coins)
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = inf
	}

	for i := 0; i <= amount; i++ {
		dp[0] = 0
	}

	for i := 1; i <= coinsCount; i++ {
		currentCoin := coins[i-1]
		for j := 1; j <= amount; j++ {
			if j >= currentCoin {
				dp[j] = minn(dp[j], dp[j-currentCoin]+1)
			}
		}
	}

	ans = dp[amount]

	if ans == inf {
		return -1
	}

	return ans
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

	coins := Deserialize[[]int](ReadLine(reader))
	amount := Deserialize[int](ReadLine(reader))
	ans := coinChange(coins, amount)

	fmt.Println("\noutput:", Serialize(ans))
}
