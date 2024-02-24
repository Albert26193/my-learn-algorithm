// Created by Albert's server at 2024/02/23 19:59
// leetgo: dev
// https://leetcode.cn/problems/simplify-path/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")

	stack := make([]string, 0)
	for _, dir := range dirs {
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if dir == "" || dir == "." {
			continue
		} else {
			stack = append(stack, dir)
		}
	}

	return "/" + strings.Join(stack, "/")
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
	path := Deserialize[string](ReadLine(in))
	ans := simplifyPath(path)

	fmt.Println("\noutput:", Serialize(ans))
}
