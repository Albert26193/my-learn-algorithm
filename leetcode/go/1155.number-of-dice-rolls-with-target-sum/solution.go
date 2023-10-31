// Created by Albert at 2023/10/24 14:06
// leetgo: 1.3.8
// https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

/*
dp[i][j] = a: choose i nums(length = i) ,and sum = j ==> a methods
dp[i][j] = dp[i - 1][j] (not choose i nth) ==> X not eligible, must choose a number

dp[i][j] = dp[i - 1][j - 1] (current num == 1)
dp[i][j] = dp[i - 1][j - 2] (current num == 2)
...
dp[i][j] = dp[i - 1][j - j] (current num == j)
==> for range
*/

func numRollsToTarget(n int, k int, target int) (ans int) {
	if target < n || target > n*k {
		return 0
	}

	const mod = 1e9 + 7

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= k && i <= target; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i; j <= target; j++ {
			for m := 1; m <= k && m <= j; m++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-m]) % mod
			}
		}
	}

	return dp[n][target]
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

	n := Deserialize[int](ReadLine(reader))
	k := Deserialize[int](ReadLine(reader))
	target := Deserialize[int](ReadLine(reader))
	ans := numRollsToTarget(n, k, target)

	fmt.Println("\noutput:", Serialize(ans))
}
