// Created by Albert's server at 2024/01/21 20:31
// leetgo: dev
// https://leetcode.cn/problems/swim-in-rising-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func swimInWater(grid [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := swimInWater(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
