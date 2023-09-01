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
func tempMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) (ans int) {
	numsLength := len(nums)
	maxStepPosition := 0

	jumpCount := 0
	end := 0
	for i := 0; i < numsLength-1; i++ {
		// 继续往下遍历，统计边界范围内，哪一格能跳得更远，每走一步就更新一次能跳跃的最远位置下标
		// 其实就是在统计下一步的最优情况
		maxStepPosition = tempMax(maxStepPosition, nums[i]+i)
		// 如果到达了边界，那么一定要跳了，下一跳的边界下标就是之前统计的最优情况maxPosition，并且步数加1
		if i == end {
			end = maxStepPosition
			jumpCount += 1
		}
	}

	return jumpCount
}

// @lc code=end

func main() {

	// nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{2, 3, 1, 1, 4}
	ans := jump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
