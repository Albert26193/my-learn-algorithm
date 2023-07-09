// Created by Albert at 2023/07/08 16:45
// leetgo: 1.3.2
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

//TODO: 结合889系列去看
// @lc code=begin
// preorder := [3, 9, 20, 15, 7].  inorder := [9, 3, 15, 20, 7]
//            3
//          9    20
//             15  7
// pre:                 3 [9] [20, 15, 7]
// breakPointIndex      ^
// in:                  [9]  3 [15 20  7]
//
//
//
func buildTree(preorder []int, inorder []int) (ans *TreeNode) {
	var getTree func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode
	getTree = func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}

		currentLength := preEnd - preStart + 1
		root := &TreeNode{Val: preorder[preStart]}
		if currentLength == 1 {
			return root
		}

		breakPointIndex := 0
		for preorder[preStart] != inorder[breakPointIndex] {
			breakPointIndex += 1
		}

		leftLength := breakPointIndex - inStart
		root.Left = getTree(preStart+1, preStart+leftLength, inStart, inStart+leftLength-1)
		root.Right = getTree(preStart+leftLength+1, preEnd, inStart+leftLength+1, inEnd)

		return root
	}
	return getTree(0, len(preorder)-1, 0, len(inorder)-1)
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// preorder := Deserialize[[]int](ReadLine(stdin))
	// inorder := Deserialize[[]int](ReadLine(stdin))

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	ans := buildTree(preorder, inorder)

	fmt.Println("\noutput:", Serialize(ans))
}
