// Created by Albert's server at 2024/02/16 16:34
// leetgo: dev
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxPathSum(root *TreeNode) (ans int) {
	maxSum := math.MinInt32

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		innerMax := left + node.Val + right
		maxSum = maxx(maxSum, innerMax)
		outputMax := node.Val + maxx(left, right)
		return maxx(outputMax, 0)
	}

	dfs(root)

	return maxSum
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
	root := Deserialize[*TreeNode](ReadLine(in))
	ans := maxPathSum(root)

	fmt.Println("\noutput:", Serialize(ans))
}
