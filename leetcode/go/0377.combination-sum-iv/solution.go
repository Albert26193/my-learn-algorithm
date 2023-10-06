// Created by Albert at 2023/10/02 16:52
// leetgo: 1.3.8
// https://leetcode.cn/problems/combination-sum-iv/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
func combinationSum4(nums []int, target int) (ans int) {
	dp := make([]int, target+1)

	dp[0] = 1
	for _, num := range nums {
		for i := num; i <= target; i++ {
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
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
	target := Deserialize[int](ReadLine(reader))
	ans := combinationSum4(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
