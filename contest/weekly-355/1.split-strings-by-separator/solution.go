// Created by Albert at 2023/09/26 14:43
// leetgo: 1.3.8
// https://leetcode.cn/problems/split-strings-by-separator/
// https://leetcode.cn/contest/weekly-contest-355/problems/split-strings-by-separator/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func splitWordsBySeparator(words []string, separator byte) (ans []string) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	words := Deserialize[[]string](ReadLine(stdin))
	separator := Deserialize[byte](ReadLine(stdin))
	ans := splitWordsBySeparator(words, separator)

	fmt.Println("\noutput:", Serialize(ans))
}
