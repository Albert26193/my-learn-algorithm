// Created by Albert's server at 2024/02/20 18:13
// leetgo: dev
// https://leetcode.cn/problems/word-break-ii/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func wordBreak(s string, wordDict []string) []string {
	n := len(s)

	ans := make([][]string, 0)
	mp := make(map[string]bool)

	for _, word := range wordDict {
		mp[word] = true
	}

	var dfs func(start int, cur []string)
	dfs = func(start int, cur []string) {
		if start == n {
			temp := make([]string, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}

		for i := start; i < n; i++ {
			curWord := s[start : i+1]
			if mp[curWord] {
				cur = append(cur, curWord)
				dfs(i+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs(0, []string{})
	res := make([]string, 0)
	for _, comb := range ans {
		str := strings.Join(comb, " ")
		res = append(res, str)
	}
	// fmt.Println(ans)
	return res
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	s := Deserialize[string](ReadLine(in))
	wordDict := Deserialize[[]string](ReadLine(in))
	ans := wordBreak(s, wordDict)

	fmt.Println("\noutput:", Serialize(ans))
}
