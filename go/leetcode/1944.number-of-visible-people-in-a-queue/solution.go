// Created by Albert's server at 2024/01/05 16:27
// leetgo: dev
// https://leetcode.cn/problems/number-of-visible-people-in-a-queue/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canSeePersonsCount(heights []int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	heights := Deserialize[[]int](ReadLine(stdin))
	ans := canSeePersonsCount(heights)

	fmt.Println("\noutput:", Serialize(ans))
}
