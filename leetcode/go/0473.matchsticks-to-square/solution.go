// Created by Albert at 2023/04/13 11:50
// https://leetcode.cn/problems/matchsticks-to-square/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func makesquare(matchsticks []int) bool {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matchsticks := Deserialize[[]int](ReadLine(stdin))
	ans := makesquare(matchsticks)
	fmt.Println("output: " + Serialize(ans))
}
