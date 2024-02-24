// Created by Albert at 2023/05/16 13:44
// https://leetcode.cn/problems/number-of-digit-one/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countDigitOne(n int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := countDigitOne(n)

	fmt.Println("\noutput:", Serialize(ans))
}
