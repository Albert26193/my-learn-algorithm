// Created by Albert at 2023/04/19 01:55
// https://leetcode.cn/problems/combination-sum-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func combinationSum2(candidates []int, target int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	candidates := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := combinationSum2(candidates, target)
	fmt.Println("output: " + Serialize(ans))
}
