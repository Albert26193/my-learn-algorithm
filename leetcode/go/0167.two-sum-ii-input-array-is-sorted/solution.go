// Created by Albert at 2023/07/08 14:22
// leetgo: 1.3.2
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// Constant-level space complexity
// binary search algorithm: from left to right, choose current number {CurNumber}, for {curNumber}, use binary search to get eligible {resetNumber}, which meets {curNumber} + {resetNumber} = target
// ----------------M**************
func binarySearchTarget(index int, numbers []int, target int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right + 1) / 2
		if numbers[index]+numbers[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if numbers[index]+numbers[right] == target {
		return right
	} else {
		return -1
	}
}

func twoSum(numbers []int, target int) (ans []int) {
	numbersLength := len(numbers)
	for i := 0; i < numbersLength; i++ {
		resetNumberIndex := binarySearchTarget(i, numbers, target)
		if resetNumberIndex != -1 {
			return []int{i + 1, resetNumberIndex + 1}
		}
	}
	return []int{-1, -1}
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// numbers := Deserialize[[]int](ReadLine(stdin))
	// target := Deserialize[int](ReadLine(stdin))

	numbers := []int{2, 7, 11, 15}
	numbers = []int{2, 3, 4}
	target := 9
	target = 6
	ans := twoSum(numbers, target)

	fmt.Println("\noutput:", Serialize(ans))
}
