// Created by Albert at 2023/07/09 12:46
// leetgo: 1.3.2
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
/*
*  postorder =           [9], [15, 7, 20], 3
*  inorder =             [9], 3, [15, 20, 7]
*  breakPointIndex = 1        ^
*
*
*
*
*
*/

func buildTree(inorder []int, postorder []int) (ans *TreeNode) {
    var getTree func(inStart int, inEnd int, postStart int, postEnd int) *TreeNode
    getTree = func(inStart int, inEnd int, postStart int, postEnd int) *TreeNode {
        if postStart > postEnd {
            return nil
        }

        currentLength := postEnd - postStart + 1
        root := &TreeNode{Val: postorder[postEnd]}
        if  currentLength == 1 {
            return root
        }

        breakPointIndex := 0
        for inorder[breakPointIndex] != postorder[postEnd] {
            breakPointIndex += 1
        }

        leftTreeLength := breakPointIndex - inStart
        root.Left = getTree(inStart, inStart + leftTreeLength - 1, postStart, postStart + leftTreeLength - 1)
        root.Right = getTree(inStart + leftTreeLength + 1, inEnd, postStart + leftTreeLength, postEnd - 1)

        return root
    }
	return getTree(0, len(inorder) - 1, 0, len(postorder) - 1)
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// inorder := Deserialize[[]int](ReadLine(stdin))
	// postorder := Deserialize[[]int](ReadLine(stdin))
    inorder := []int {9, 3, 15, 20, 7}
    postorder := []int {9, 15, 7, 20, 3}

	ans := buildTree(inorder, postorder)

	fmt.Println("\noutput:", Serialize(ans))
}
