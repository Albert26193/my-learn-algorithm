// Created by Albert at 2023/04/15 14:13
// https://leetcode.cn/problems/flower-planting-with-no-adjacent/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 1. the graph may not be connected
// 2. must be a solution
func gardenNoAdj(n int, paths [][]int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	paths := Deserialize[[][]int](ReadLine(stdin))
	ans := gardenNoAdj(n, paths)
	fmt.Println("output: " + Serialize(ans))
}
