// Created by Albert's server at 2024/02/21 02:44
// leetgo: dev
// https://leetcode.cn/problems/jump-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canJump(nums []int) bool {
	maxReach := 0
	for i, num := range nums {
		// 如果当前位置超过了之前计算的最远到达点，说明这个位置不可达，返回false
		if i > maxReach {
			return false
		}
		// 更新最远到达点
		if i+num > maxReach {
			maxReach = i + num
		}
		// 如果最远到达点已经超过或达到数组的最后一个位置，返回true
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	nums := Deserialize[[]int](ReadLine(in))
	ans := canJump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
