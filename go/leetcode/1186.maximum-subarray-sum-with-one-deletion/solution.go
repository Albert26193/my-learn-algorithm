// Created by Albert at 2023/06/27 08:25
// leetgo: 1.3.2
// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumSum(arr []int) (ans int) {
	// dp[i][0] = max(dp[i - 1][0], 0) + arr[i]
	// dp[i][1] = max(dp[i - 1][1] + arr[i], dp[i - 1][0])
	// dp[i-1][0] ==> dp0.  dp[i - 1][1] ==> dp1
	dp0, dp1 := arr[0], 0
	ans = arr[0]
	for i := 1; i < len(arr); i++ {
		dp0, dp1 = max(dp0, 0)+arr[i], max(dp1+arr[i], dp0)
		ans = max(ans, max(dp0, dp1))
	}

	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// arr := Deserialize[[]int](ReadLine(stdin))

	arr := []int{1, -2, 0, 3}
	ans := maximumSum(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
