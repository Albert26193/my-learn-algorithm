// Created by Albert at 2023/04/05 23:01
// https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

//TODO
func findTheLongestBalancedSubstring(s string) (ans int) {
	hasBegin := false
	hasOne := false
	rec := 0
	beginIndex := 0
	tempAns := -1
	for index, value := range s {
		if value == '0' {
			if !hasBegin && !hasOne {
				beginIndex = index
			}
			hasBegin = true
			rec -= 1
		}

		if !hasBegin {
			continue
		}

		if value == '1' {
			rec += 1
			hasOne = true
		}

		if rec <= 0 && hasOne {
			tempAns = index - beginIndex + 1 + rec
			if tempAns > ans {
				ans = tempAns
			}
		}

		if rec == 0 {
			hasOne = false
			hasBegin = false
		}
	}

	return ans
}

// @lc code=end

func main() {
	//stdin := bufio.NewReader(os.Stdin)
	// s := Deserialize[string](ReadLine(stdin))
	s := "010101011110000001"
	ans := findTheLongestBalancedSubstring(s)
	fmt.Println("output: " + Serialize(ans))
}
