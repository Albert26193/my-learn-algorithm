// Created by Albert's server at 2024/05/16 14:40
// leetgo: dev
// https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numberOfWeeks(milestones []int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	milestones := Deserialize[[]int](ReadLine(stdin))
	ans := numberOfWeeks(milestones)

	fmt.Println("\noutput:", Serialize(ans))
}
