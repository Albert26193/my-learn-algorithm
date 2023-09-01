// Created by Albert at 2023/09/01 18:00
// leetgo: 1.3.7
// https://leetcode.cn/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
