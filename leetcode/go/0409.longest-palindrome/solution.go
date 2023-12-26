// Created by Albert's server at 2023/12/26 15:07
// leetgo: dev
// https://leetcode.cn/problems/longest-palindrome/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func getcode(ch byte) int {
	if 'a' <= ch && ch <= 'z' {
		return int(ch - 'a')
	}

	if 'A' <= ch && ch <= 'Z' {
		return int(ch-'A') + 26
	}
	return -1
}

func longestPalindrome(s string) (ans int) {
	n := len(s)

	recSize := 52
	rec := make([]int, recSize)

	for i := 0; i < n; i++ {
		rec[getcode(s[i])] += 1
		// fmt.Println(getcode(s[i]))
	}

	flag := 0
	for i := 0; i < recSize; i++ {
		if rec[i]&1 == 1 {
			flag = 1
			ans += rec[i] - 1
		} else {
			ans += rec[i]
		}
	}

	return ans + flag
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
