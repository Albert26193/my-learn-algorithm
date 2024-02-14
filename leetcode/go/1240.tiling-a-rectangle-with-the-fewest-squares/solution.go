// Created by Albert's server at 2024/01/06 19:35
// leetgo: dev
// https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func tilingRectangle(n int, m int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	ans := tilingRectangle(n, m)

	fmt.Println("\noutput:", Serialize(ans))
}
