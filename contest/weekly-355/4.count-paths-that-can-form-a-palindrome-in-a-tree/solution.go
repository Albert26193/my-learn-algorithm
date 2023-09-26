// Created by Albert at 2023/09/26 14:43
// leetgo: 1.3.8
// https://leetcode.cn/problems/count-paths-that-can-form-a-palindrome-in-a-tree/
// https://leetcode.cn/contest/weekly-contest-355/problems/count-paths-that-can-form-a-palindrome-in-a-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countPalindromePaths(parent []int, s string) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	parent := Deserialize[[]int](ReadLine(stdin))
	s := Deserialize[string](ReadLine(stdin))
	ans := countPalindromePaths(parent, s)

	fmt.Println("\noutput:", Serialize(ans))
}
