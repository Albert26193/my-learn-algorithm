// Created by Albert at 2023/09/18 10:08
// leetgo: 1.3.8
// https://leetcode.cn/problems/house-robber-iv/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 1. 二分答案 + DP;   2. 二分答案 + 贪心

// 2. 二分答案 + 贪心
func minCapability(nums []int, k int) (ans int) {
	left, right := 0, int(1e9)
	for left < right {
		mid := (left + right) / 2
		count := 0
		visited := false

		for _, x := range nums {
			if x <= mid && !visited {
				count += 1
				visited = true
			} else {
				visited = false
			}
		}

		// 如果符合条件
		if count >= k {
			right = mid
		} else {
			// 如果不符合条件
			left = mid + 1
		}
	}

	return left
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
	k := Deserialize[int](ReadLine(reader))

	ans := minCapability(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
