// Created by Albert at 2023/11/13 00:37
// leetgo: 1.3.8
// https://leetcode.cn/problems/letter-case-permutation/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: 1. BFS;  2. backTrack(2 methods);  3.bitmap mask
func letterCasePermutation(s string) (ans []string) {
	var dfs func(int, []byte)
	dfs = func(index int, currentString []byte) {
		for index < len(currentString) && unicode.IsDigit(rune(currentString[index])) {
			index += 1
		}

		if index == len(currentString) {
			ans = append(ans, string(currentString))
			return
		}

		dfs(index+1, currentString)

		(currentString)[index] ^= 32

		dfs(index+1, currentString)
	}

	dfs(0, []byte(s))

	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	s := Deserialize[string](ReadLine(reader))
	ans := letterCasePermutation(s)

	fmt.Println("\noutput:", Serialize(ans))
}
