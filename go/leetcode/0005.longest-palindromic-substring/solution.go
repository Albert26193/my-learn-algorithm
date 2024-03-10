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

// f[i][j] ==> [i, j] is Palindrome ?
// f[i][j] = f[i+1][j-1] && s[i == s[j]
func longestPalindrome(s string) (ans string) {
	n := len(s)

	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for begin := 0; begin+length-1 <= n-1; begin++ {
			end := begin + length - 1

			if s[begin] != s[end] {
				f[begin][end] = false
				continue
			}

			if length <= 3 {
				f[begin][end] = true
				ans = s[begin : end+1]
				continue
			}

			f[begin][end] = f[begin+1][end-1]

			if f[begin][end] {
				ans = s[begin : end+1]
			}
		}
	}
	return
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
