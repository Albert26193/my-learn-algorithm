/*
 * @lc app=leetcode.cn id=831 lang=golang
 *
 * [831] 隐藏个人信息
 */
package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func convertPhone(s string) (ans string) {
	// get suffix
	processedString := strings.ReplaceAll(s, "-", "")
	processedString = strings.ReplaceAll(processedString, "+", "")
	processedString = strings.ReplaceAll(processedString, "(", "")
	processedString = strings.ReplaceAll(processedString, ")", "")
	processedString = strings.ReplaceAll(processedString, " ", "")
	suffix := processedString[len(processedString) - 4: ]
	suffix = "***-***-" + suffix	

	// get prefix
	if len(processedString) < 10 {
		return "error111"
	}
	prefixLength := len(processedString) - 10
	if prefixLength == 0 {
		return suffix
	} else if prefixLength > 3 {
		return "error"
	} else {
		return "+" + strings.Repeat("*", prefixLength) + "-" + suffix
	}
}

func convertEmail(s string) (ans string) {
	startChar := s[0]
	symbolIndex := strings.Index(s, "@")
	if symbolIndex == -1 {
		return "error"
	}
	endChar := s[symbolIndex - 1]
	prefixString := string(startChar) + "*****" + string(endChar)
	fullString := prefixString + s[symbolIndex:]
	return strings.ToLower(fullString)
}

func maskPII(s string) (ans string) {
	if strings.Contains(s, "@") {
		ans = convertEmail(s)
	} else {
		ans = convertPhone(s)
	}

	return ans
}
// @lc code=end

func main() {
	ans1 := maskPII("LeetCode@LeetCode.com")
	ans2 := maskPII("AB@qq.com")
	ans3 := maskPII("86-(10)12345678")
	ans4 := maskPII("1(234)567-890")
	fmt.Println(ans1, ans2, ans3, ans4)
}