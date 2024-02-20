// Created by Albert's server at 2024/02/16 21:12
// leetgo: dev
// https://leetcode.cn/problems/sum-root-to-leaf-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode, preSum int) int
	dfs = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0
		}

		if root.Left == nil && root.Right == nil {
			return preSum * 10 + root.Val
		}
		res := 0
		if root.Left != nil {
			cur := preSum*10 + root.Val
			res += dfs(root.Left, cur)
		}
		if root.Right != nil {
			cur := preSum*10 + root.Val
			res += dfs(root.Right, cur)
		}

		return res		
	}

	return dfs(root, 0)
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
	ans := sumNumbers(root)

	fmt.Println("\noutput:", Serialize(ans))
}
