// Created by Albert at 2023/10/02 16:50
// leetgo: 1.3.8
// https://leetcode.cn/problems/perfect-squares/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
func minn(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

func numSquares(n int) (ans int) {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 1; i <= n; i++ {
		minNum := 10001
		for j := 1; j*j <= i; j++ {
			minNum = minn(minNum, dp[i-j*j])
		}
		dp[i] = minNum + 1
	}

	return dp[n]
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
	ans := numSquares(n)

	fmt.Println("\noutput:", Serialize(ans))
}
