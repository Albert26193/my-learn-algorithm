// Created by Albert at 2023/04/19 01:52
// https://leetcode.cn/problems/permutations-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func permuteUnique(nums []int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permuteUnique(nums)
	fmt.Println("output: " + Serialize(ans))
}
