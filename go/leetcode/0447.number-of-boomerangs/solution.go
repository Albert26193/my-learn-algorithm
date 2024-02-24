// Created by Albert at 2023/10/09 08:45
// leetgo: 1.3.8
// https://leetcode.cn/problems/number-of-boomerangs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO:
func numberOfBoomerangs(points [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	points := Deserialize[[][]int](ReadLine(stdin))
	ans := numberOfBoomerangs(points)

	fmt.Println("\noutput:", Serialize(ans))
}
