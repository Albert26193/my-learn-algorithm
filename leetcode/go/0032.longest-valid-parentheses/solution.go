// Created by Albert's server at 2024/02/14 18:07
// leetgo: dev
// https://leetcode.cn/problems/longest-valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestValidParentheses(s string) (ans int) {
	n := len(s)

	// f[i] ==> end with s[i], max len is f[i]
	f := make([]int, n+1)
	f[0] = 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					f[i] = f[i-2] + 2
				} else {
					f[i] = 2
				}
			} else if i-f[i-1] >= 1 && s[i-f[i-1]-1] == '(' {
				if i-f[i-1] >= 2 {
					f[i] = f[i-1] + f[i-f[i-1]-2] + 2
				} else {
					f[i] = f[i-1] + 2
				}
			}

			ans = maxx(ans, f[i])
		}
	} 
	return ans
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

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
	ans := longestValidParentheses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
