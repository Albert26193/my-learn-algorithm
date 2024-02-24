// Created by Albert at 2023/06/26 07:03
// leetgo: 1.3.2
// https://leetcode.cn/problems/decode-ways/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

//TODO: 抄的别人题解，有空回头看
// @lc code=begin
func numDecodings(s string) int {
	originLength := len(s)
	s = " " + s
	dp := make([]int, originLength+5)
	dp[0] = 1
	for i := 1; i <= originLength; i++ {
		firstNum := s[i] - '0'
		secondNum := (s[i-1]-'0')*10 + (s[i] - '0')
		if 1 <= firstNum && firstNum <= 9 {
			dp[i] = dp[i-1]
		}

		if 10 <= secondNum && secondNum <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[originLength]
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))

	s := "226"
	// s = "06"
	ans := numDecodings(s)

	fmt.Println("\noutput:", Serialize(ans))
}
