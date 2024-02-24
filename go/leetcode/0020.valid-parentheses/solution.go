// Created by Albert at 2023/09/14 21:20
// leetgo: 1.3.8
// https://leetcode.cn/problems/valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValid(s string) bool {
	stack := []rune{}

	pairMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				if stack[len(stack)-1] != pairMap[ch] {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return len(stack) == 0
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
	ans := isValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
