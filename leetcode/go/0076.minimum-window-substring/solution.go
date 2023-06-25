// Created by Albert at 2023/06/25 16:10
// leetgo: 1.3.2
// https://leetcode.cn/problems/minimum-window-substring/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type Position struct {
	minLength int
	subString string
}

func checkWindowStatus(recordCounts []int, patternCounts []int) bool {
	for i := 0; i < len(patternCounts); i++ {
		if recordCounts[i] < patternCounts[i] {
			return false
		}
	}

	return true
}

func minWindow(s string, t string) string {
	minLengthPosition := Position{
		minLength: 0x3f3f3f3f,
		subString: "",
	}

	recordCounts := make([]int, 126)
	patternCounts := make([]int, 126)
	originStringLength := len(s)
	startIndex, endIndex := 0, 0

	for index := range t {
		tCharCode := t[index] - 'A'
		patternCounts[tCharCode] += 1
	}

	for endIndex = 0; endIndex < originStringLength; endIndex += 1 {
		endCharCode := s[endIndex] - 'A'
		recordCounts[endCharCode] += 1

		for (checkWindowStatus(recordCounts, patternCounts)) && (startIndex <= endIndex) {
			eligibleLength := (endIndex - startIndex + 1)
			if eligibleLength < minLengthPosition.minLength {
				minLengthPosition.minLength = eligibleLength
				minLengthPosition.subString = s[startIndex : endIndex+1]
			}

			startCharCode := s[startIndex] - 'A'
			recordCounts[startCharCode] -= 1
			startIndex += 1
		}

	}

	return minLengthPosition.subString
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))
	// t := Deserialize[string](ReadLine(stdin))

	s := "ADOBECODEBANC"
	t := "ABC"

	s = "A"
	t = "A"
	ans := minWindow(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
