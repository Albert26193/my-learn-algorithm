// Created by Albert at 2023/04/18 23:26
// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

//TODO
// @lc code=begin

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// minVal: minVal in both left children and right children
// maxVal: maxVal in both left children and right children
// 1. largest distance is abs(root.Val - someChildNode.Val)
// 2. largest distance is abs(leftSonVal - someLeftChildNode.Val)
// 3. largest distance is abs(rightSonVal - someRightChildNode.Val)
func dfs(root *TreeNode, minVal int, maxVal int) int {
	if root == nil {
		return 0
	}

	diff := max(abs(root.Val-minVal), abs(root.Val-maxVal))
	minVal = min(minVal, root.Val)
	maxVal = max(maxVal, root.Val)
	diff = max(diff, dfs(root.Left, minVal, maxVal))
	diff = max(diff, dfs(root.Right, minVal, maxVal))
	return diff
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	return dfs(root, root.Val, root.Val)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := maxAncestorDiff(root)
	fmt.Println("output: " + Serialize(ans))
}
