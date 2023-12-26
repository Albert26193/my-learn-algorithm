// Created by Albert's server at 2023/12/26 15:24
// leetgo: dev
// https://leetcode.cn/problems/longest-palindromic-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// f[i][j] bool ==> [i, ..., j] is true or false
// f[i][j] : true : f[i-1][j-1] && s[i] == s[j]
// f[i][j] : false: otherwise

func longestPalindrome(s string) (ans string) {
	n := len(s)

	if n < 2 {
		return s
	}

	ans = s[0:1]
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
	}

	// boundary
	for i := 0; i < n; i++ {
		f[i][i] = true
	}

	for len := 2; len <= n; len++ {
		for start := 0; start+len-1 < n; start++ {
			end := start + len - 1

			if s[start] != s[end] {
				f[start][end] = false
				continue
			}

			if len <= 3 {
				f[start][end] = true
				ans = s[start : end+1]
				continue
			}

			f[start][end] = f[start+1][end-1]
			if f[start][end] {
				// fmt.Println(start, end)
				ans = s[start : end+1]
			}
		}
	}

	return ans
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
	ans := longestPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
