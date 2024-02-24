// Created by Albert at 2023/09/14 22:02
// leetgo: 1.3.8
// https://leetcode.cn/problems/valid-parenthesis-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 1. 贪心   2. 栈  3. DP
func checkValidString(s string) bool {
	count := 0
    for 
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := checkValidString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
