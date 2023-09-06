// Created by Albert at 2023/09/06 14:50
// leetgo: 1.3.7
// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// TODO: 此题还有另外的办法
// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/solutions/2428724/liang-chong-di-gui-si-lu-pythonjavacgojs-xxnk/
// @lc code=begin

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, ans := findDeepestNode(root)
	return ans
}

func findDeepestNode(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	// why error?
	if root.Left == nil && root.Right == nil {
		return 1, root
	}

	depthLeft, leftLcaNode := findDeepestNode(root.Left)
	depthRight, rightLcaNode := findDeepestNode(root.Right)

	if depthLeft < depthRight {
		return depthRight + 1, rightLcaNode
	} else if depthLeft > depthRight {
		return depthLeft + 1, leftLcaNode
	} else {
		return depthLeft + 1, root
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := lcaDeepestLeaves(root)

	fmt.Println("\noutput:", Serialize(ans))
}
