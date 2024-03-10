// Created by Albert's server at 2024/02/14 14:07
// leetgo: dev
// https://leetcode.cn/problems/palindrome-partitioning/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func partition(s string) (ans [][]string) {
	var dfs func(cur []string, start int)

	n := len(s)
	// start:
	dfs = func(cur []string, start int) {
		if start == n {
			ans = append(ans, append([]string{}, cur...))
			return
		}

		// enum end position
		for i := start; i < n; i++ {
			comb := s[start : i+1]

			if isPali(s, start, i) {
				cur = append(cur, comb)
				dfs(cur, i+1)
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs([]string{}, 0)
	return
}

func isPali(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
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
	ans := partition(s)

	fmt.Println("\noutput:", Serialize(ans))
}
