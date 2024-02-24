// Created by Albert at 2023/06/24 22:05
// https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func getCountOfRecord(record []int) (eligibleCount int) {
	recordLength := len(record)
	eligibleCount = 0

	if recordLength != 126 {
		return -1
	}

	for i := 0; i < recordLength; i++ {
		if record[i] > 0 {
			eligibleCount += 1
		}
	}
	return eligibleCount
}

func calculateMath(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func lengthOfLongestSubstringTwoDistinct(s string) (ans int) {
	record := make([]int, 126)
	left, right := 0, 0
	originStringLength := len(s)
	maxSubStringLength := -1

	for right = 0; right < originStringLength; right += 1 {
		rightOffset := s[right] - 'A'
		// fmt.Println(rightOffset)
		record[rightOffset] += 1

		for getCountOfRecord(record) > 2 && left < right {
			leftOffset := s[left] - 'A'
			record[leftOffset] -= 1
			left += 1
		}

		maxSubStringLength = calculateMath(maxSubStringLength, right-left+1)
	}
	return maxSubStringLength
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))

	s := "ccaabbb"
	s = "abZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZYX"
	ans := lengthOfLongestSubstringTwoDistinct(s)

	fmt.Println("\noutput:", Serialize(ans))
}
