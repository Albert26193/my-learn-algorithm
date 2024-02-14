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
	// memo ==> 1: true
	// memo ==> 2: false
	// memo ==> 0: not visited
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}

	var dfs func(cur []string, start int)
	dfs = func(cur []string, start int) {
		if start == len(s) {
			temp := make([]string, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}

		for i := start; i < len(s); i++ {
			if memo[start][i] == 2 {
				continue
			}

			if memo[start][i] == 1 || isPali(s, start, i, memo) {
				cur = append(cur, s[start:i+1])
				dfs(cur, i+1)
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs([]string{}, 0)
	return
}

func isPali(s string, l int, r int, memo [][]int) bool {
	for l < r {
		if s[l] != s[r] {
			memo[l][r] = 2
			return false
		}
		l++
		r--
	}

	memo[l][r] = 1
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
