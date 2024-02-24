// Created by Albert's server at 2024/01/22 01:14
// leetgo: dev
// https://leetcode.cn/problems/most-profit-assigning-work/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	difficulty := Deserialize[[]int](ReadLine(stdin))
	profit := Deserialize[[]int](ReadLine(stdin))
	worker := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfitAssignment(difficulty, profit, worker)

	fmt.Println("\noutput:", Serialize(ans))
}
