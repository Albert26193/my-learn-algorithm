// Created by Albert at 2023/09/27 12:31
// leetgo: 1.3.8
// https://leetcode.cn/problems/greatest-sum-divisible-by-three/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
func maxx(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSumDivThree(nums []int) (ans int) {
	numsLength := len(nums)
	dp := make([][3]int, numsLength+1)

	// dp[i][j] ==> nums[0]...nums[i - 1] ==> maxSum % 3 == j
	dp[0][0] = 0
	dp[0][1] = -1000
	dp[0][2] = -1000

	for i := 1; i <= numsLength; i++ {
		for j := 0; j <= 2; j++ {
			dp[i][j] = maxx(dp[i-1][j], dp[i-1][(j+nums[i-1]+3)%3]+nums[i-1])
		}
	}

	return dp[numsLength][0]
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

	nums := Deserialize[[]int](ReadLine(reader))
	ans := maxSumDivThree(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
