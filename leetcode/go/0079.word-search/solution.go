// Created by Albert at 2023/11/23 16:22
// leetgo: 1.3.8
// https://leetcode.cn/problems/word-search/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func exist(board [][]byte, word string) bool {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	board := Deserialize[[][]byte](ReadLine(stdin))
	word := Deserialize[string](ReadLine(stdin))
	ans := exist(board, word)

	fmt.Println("\noutput:", Serialize(ans))
}
