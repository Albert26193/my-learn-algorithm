// Created by Albert at 2023/09/15 12:35
// leetgo: 1.3.8
// https://leetcode.cn/problems/generate-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateParenthesis(n int) (ans []string) {
	var dfs func(int, int, int, string)
	dfs = func(depth int, left int, right int, combination string) {

		if left == n && right == n {
			ans = append(ans, combination)
		}

		if depth >= 2*n {
			return
		}

		if left < n {
			combination += "("
			dfs(depth+1, left+1, right, combination)
			combination = combination[:len(combination)-1]
		}

		if right < left && right < n {
			combination += ")"
			dfs(depth+1, left, right+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	dfs(0, 0, 0, "")
	return
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

	stdin := bufio.NewReader(reader)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateParenthesis(n)

	fmt.Println("\noutput:", Serialize(ans))
}
