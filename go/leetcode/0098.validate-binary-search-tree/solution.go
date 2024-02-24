// Created by Albert's server at 2024/01/06 21:11
// leetgo: dev
// https://leetcode.cn/problems/validate-binary-search-tree/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower int, higher int) bool
	dfs = func(node *TreeNode, lower int, higher int) bool {
		if node == nil {
			return true
		}

		if node.Val <= lower || node.Val >= higher {
			return false
		}

		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, higher)
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
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
	root := Deserialize[*TreeNode](ReadLine(in))
	ans := isValidBST(root)

	fmt.Println("\noutput:", Serialize(ans))
}
