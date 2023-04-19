// Created by Albert at 2023/04/13 11:49
// https://leetcode.cn/problems/word-squares/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func wordSquares(words []string) (ans [][]string) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	words := Deserialize[[]string](ReadLine(stdin))
	ans := wordSquares(words)
	fmt.Println("output: " + Serialize(ans))
}
