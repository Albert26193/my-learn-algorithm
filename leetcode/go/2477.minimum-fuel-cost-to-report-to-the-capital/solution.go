// Created by Albert's server at 2023/12/05 12:23
// leetgo: dev
// https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumFuelCost(roads [][]int, seats int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	roads := Deserialize[[][]int](ReadLine(stdin))
	seats := Deserialize[int](ReadLine(stdin))
	ans := minimumFuelCost(roads, seats)

	fmt.Println("\noutput:", Serialize(ans))
}
