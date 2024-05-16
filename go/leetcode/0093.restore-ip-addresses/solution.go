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
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return 0 <= num && num <= 255
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
	var dfs func(index int, cur []string)
	dfs = func(index int, cur []string) {
		if len(cur) == 4 {
			if index == len(s) {
				ans = append(ans, getString(cur))
			}
			return
		}

		for l := 1; l <= 3; l++ {
			if index+l > len(s) {
				continue
			}
			str := s[index : index+l]
			if !check(str) {
				continue
			}
			cur = append(cur, str)
			dfs(index+l, cur)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, []string{})
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
	ans := restoreIpAddresses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
