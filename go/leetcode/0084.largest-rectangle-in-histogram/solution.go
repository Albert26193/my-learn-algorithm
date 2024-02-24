// Created by Albert at 2023/06/29 13:55
// leetgo: 1.3.2
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
//TODO: Monotonic stack
func largestRectangleArea(heights []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	heights := Deserialize[[]int](ReadLine(stdin))
	ans := largestRectangleArea(heights)

	fmt.Println("\noutput:", Serialize(ans))
}
