// Created by Albert at 2023/06/30 13:14
// leetgo: 1.3.2
// https://leetcode.cn/problems/sum-of-subarray-ranges/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findLeftAndRightBoundary(nums []int) (boundaryRecord [][2]int) {
	//TODO
	numsLength := len(nums)
	boundaryRecord = make([][2]int, numsLength)

	leftMonotonicStack := make([]int, 0)

	for i := 0; i < numsLength; i++ {
		currentPushNum := nums[i]
		for len(leftMonotonicStack) > 0 && (currentPushNum >= nums[leftMonotonicStack[len(leftMonotonicStack)-1]]) {
			leftMonotonicStack = leftMonotonicStack[0 : len(leftMonotonicStack)-1]
		}

		// currentPushNum < stackTopValue
		if len(leftMonotonicStack) > 0 {
			boundaryRecord[i][0] = leftMonotonicStack[len(leftMonotonicStack)-1]
		} else {
			boundaryRecord[i][1] = -1
		}

		leftMonotonicStack = append(leftMonotonicStack, i)
	}

	return
}
func subArrayRanges(nums []int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := subArrayRanges(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
