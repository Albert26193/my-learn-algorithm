// Created by Albert at 2023/07/09 16:43
// leetgo: 1.3.2
// https://leetcode.cn/problems/3sum/

package main

import (
	"fmt"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 需要二刷一下，题目还是非常不错的，此外，检查数组范围合法性，请务必独立出函数
func checkRange(left int, right int, nums []int) bool {
	leftBoundary := 0
	rightBoundary := len(nums) - 1
	if !(left < right) {
		return false
	}

	if !(leftBoundary <= left && left <= rightBoundary) {
		return false
	}

	if !(leftBoundary <= right && right <= rightBoundary) {
		return false
	}

	return true
}

// why it is wrong?
// for one index, must Get ALL of possible pairs
// otherwise, it will miss some answers because de-duplication
func calTwoSumIndexWrong(nums []int, left int, right int, targetSum int) []int {
	leftPointer := left
	rightPointer := right
	errorPair := []int{0x3f3f3f3f, 0x3f3f3f3f}

	if !(checkRange(leftPointer, rightPointer, nums)) {
		return errorPair
	}

	for checkRange(leftPointer, rightPointer, nums) && (nums[leftPointer]+nums[rightPointer] < targetSum) {
		leftPointer += 1
	}

	for checkRange(leftPointer, rightPointer, nums) && (nums[leftPointer]+nums[rightPointer] > targetSum) {
		rightPointer -= 1
	}

	if checkRange(leftPointer, rightPointer, nums) && nums[leftPointer]+nums[rightPointer] == targetSum {
		return []int{nums[leftPointer], nums[rightPointer]}
	} else {
		return errorPair
	}
}

func calTwoSumIndex(nums []int, left int, right int, targetSum int, ans *[][]int) {
	leftPointer := left
	rightPointer := right

	if !(checkRange(leftPointer, rightPointer, nums)) {
		return
	}

	for checkRange(leftPointer, rightPointer, nums) && leftPointer < rightPointer {
		if nums[leftPointer]+nums[rightPointer] < targetSum {
			leftPointer += 1
		} else if nums[leftPointer]+nums[rightPointer] > targetSum {
			rightPointer -= 1
		} else {
			*ans = append(*ans, []int{-targetSum, nums[leftPointer], nums[rightPointer]})
			leftPointer += 1
			rightPointer -= 1

			for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer-1] {
				leftPointer += 1
			}

			for leftPointer < rightPointer && nums[rightPointer] == nums[rightPointer+1] {
				rightPointer -= 1
			}

		}
	}

	return
}

// -4 -1 -1 0 1 2
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for i, num := range nums {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}

		calTwoSumIndex(nums, i+1, len(nums)-1, -num, &ans)
	}
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums = []int{0, 0, 0, 0}
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
