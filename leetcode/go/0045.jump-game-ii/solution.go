// Created by Albert at 2023/09/01 18:02
// leetgo: 1.3.7
// https://leetcode.cn/problems/jump-game-ii/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// TODO
// @lc code=begin
func jump(nums []int) int {
	jumps := 0      // 跳跃次数
	currentEnd := 0 // 当前跳跃可以到达的最远位置
	farthest := 0   // 所有能跳到的位置中，最远的那个位置

	// 遍历数组，但不包括最后一个元素，因为到达最后一个元素时不需要再跳跃
	for i := 0; i < len(nums)-1; i++ {
		// 更新能跳到的最远位置
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// 如果到达了当前跳跃可以到达的最远位置
		// 需要进行下一次跳跃，更新当前跳跃的最远位置为之前计算的最远位置
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

// @lc code=end

func main() {

	// nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{2, 3, 1, 1, 4}
	ans := jump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
