// Created by Albert's server at 2024/01/14 14:59
// leetgo: dev
// https://leetcode.cn/problems/restore-ip-addresses/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func check(s string) bool {
	if len(s) >= 4 {
		return false
	}
	if len(s) >= 2 && s[0] == '0' {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if num < 0 || num > 255 {
		return false
	}

	return true
}

func getString(ss []string) string {
	res := ""
	for _, s := range ss {
		res += s
		res += "."
	}

	return res[:len(res)-1]
}

func restoreIpAddresses(s string) (ans []string) {
	n := len(s)

	ans = make([]string, 0)
	var dfs func(index int, cur []string)
	dfs = func(index int, cur []string) {
		if index == n {
			if len(cur) == 4 {
				ans = append(ans, getString(cur))
			}
		}

		if len(cur) >= 4 {
			return
		}

		// enum length
		for l := 1; l <= 3; l++ {
			if l+index > n {
				continue
			}

			next := s[index : l+index]
			if !check(next) {
				continue
			}

			cur = append(cur, next)
			dfs(index+l, cur)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, []string{})
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
	ans := restoreIpAddresses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
