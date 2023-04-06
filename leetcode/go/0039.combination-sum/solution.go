// Created by Albert at 2023/04/06 14:36
// https://leetcode.cn/problems/combination-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO

func combinationSum(candidates []int, target int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	candidates := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := combinationSum(candidates, target)
	fmt.Println("output: " + Serialize(ans))
}
