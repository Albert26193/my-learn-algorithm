// Created by Albert at 2023/06/26 07:03
// leetgo: 1.3.2
// https://leetcode.cn/problems/decode-ways/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

//TODO
// @lc code=begin
func stringToNum(subString string) int {
	if len(subString) == 1 {
		return int(subString[0] - '0')
	}

	tenPosition := int(subString[0] - '0')
	singlePosition := int(subString[1] - '0')
	res := 10*tenPosition + singlePosition
	if res < 10 {
		return 1e5
	}

	return res
}

func numDecodings(s string) (ans int) {
	originLength := len(s)
	dp := make([]int, originLength+5)

	// dp[i]: 字符串的前i位 可以做 dp[i] 种拆解
	dp[0] = 0

	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= originLength; i += 1 {
		if stringToNum(s[i-2:i]) <= 26 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}

		if s[i-1] == '0' {
			dp[i] -= 1
		}
	}

	if dp[originLength] < 0 {
		return 0
	}
	return dp[originLength]
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))

	s := "10"
	// s = "06"
	ans := numDecodings(s)

	fmt.Println("\noutput:", Serialize(ans))
}
