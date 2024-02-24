// Created by Albert at 2023/04/19 01:57
// https://leetcode.cn/problems/combination-sum-iii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func combinationSum3(k int, n int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	k := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := combinationSum3(k, n)
	fmt.Println("output: " + Serialize(ans))
}
