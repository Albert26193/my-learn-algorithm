// Created by Albert at 2023/04/06 18:17
// https://leetcode.cn/problems/convert-to-base-2/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func baseNeg2(n int) string {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := baseNeg2(n)
	fmt.Println("output: " + Serialize(ans))
}
