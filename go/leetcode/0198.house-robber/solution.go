// Created by Albert at 2023/09/16 21:47
// leetgo: 1.3.8
// https://leetcode.cn/problems/house-robber/

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

func rob(nums []int) (ans int) {
	numsLength := len(nums)
	dp := make([]int, numsLength+1)

	dp[0] = 0
	dp[1] = nums[0]

	if numsLength <= 1 {
		return dp[numsLength]
	}

	dp[2] = maxx(nums[0], nums[1])

	// space optimization
	prev := 0
	current := 0
	for i := range nums {
		temp := maxx(current, prev+nums[i])
		prev = current
		current = temp
	}
	ans = current
	return
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

	stdin := bufio.NewReader(reader)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := rob(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
