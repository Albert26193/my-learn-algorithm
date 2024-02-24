// Created by Albert at 2023/09/06 14:04
// leetgo: 1.3.7
// https://leetcode.cn/problems/form-smallest-number-from-two-digit-arrays/

package main

import (
	// "bufio"
	"fmt"
	"math"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findMinNumber(nums []int) int {
	minNumber := math.MaxInt64

	for _, num := range nums {
		minNumber = min(minNumber, num)
	}

	return minNumber
}

func hasSameNumber(nums1 []int, nums2 []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums1 {
		numMap[num] = true
	}

	minNum := math.MaxInt64
	for _, num := range nums2 {
		if _, isSame := numMap[num]; isSame {
			minNum = min(minNum, num)
		}
	}

	return minNum
}

func minNumber(nums1 []int, nums2 []int) (ans int) {
	sameNumber := hasSameNumber(nums1, nums2)
	if sameNumber != math.MaxInt64 {
		return sameNumber
	}

	minNumber1 := findMinNumber(nums1)
	minNumber2 := findMinNumber(nums2)

	if minNumber1 < minNumber2 {
		return minNumber1*10 + minNumber2
	} else {
		return minNumber2*10 + minNumber1
	}
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// nums1 := Deserialize[[]int](ReadLine(stdin))
	// nums2 := Deserialize[[]int](ReadLine(stdin))

	nums1 := []int{4, 1, 3}
	nums2 := []int{5, 7}
	ans := minNumber(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
