// Created by Albert at 2023/10/30 22:37
// leetgo: 1.3.8
// https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// dp[i][j]: for s1: take 0...i-1; for s2: take 0...j-1 ==> cost
// dp[i][j] = dp[i - 1][j - 1]    if s1[i - 1] == s2[j - 1]
// dp[i][j] = dp[i - 1][j - 1] + minCost(s1[i - 1], s2[j - 1])

// dp[i][0] = dp[i - 1][0] + s1[i]
// dp[0][i] = dp[0][i = 1] + s2[i]
func minCost(a ...int) int {
	res := a[0]
	for _, num := range a {
		if num < res {
			res = num
		}
	}

	return res
}

func minimumDeleteSum(s1 string, s2 string) (ans int) {
	s1Length := len(s1)
	s2Length := len(s2)

	dp := make([][]int, s1Length+1)
	for i := 0; i <= s1Length; i++ {
		dp[i] = make([]int, s2Length+1)
	}

	for i := 1; i <= s1Length; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}

	for i := 1; i <= s2Length; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= s1Length; i++ {
		for j := 1; j <= s2Length; j++ {
			c1, c2 := s1[i-1], s2[j-1]
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minCost(dp[i-1][j]+int(c1), dp[i][j-1]+int(c2))
			}
		}
	}

	return dp[s1Length][s2Length]
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

	s1 := Deserialize[string](ReadLine(reader))
	s2 := Deserialize[string](ReadLine(reader))
	// fmt.Println(s1, s2)
	ans := minimumDeleteSum(s1, s2)

	fmt.Println("\noutput:", Serialize(ans))
}
