// Created by Albert at 2023/07/08 15:25
// leetgo: 1.3.2
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
//TODO: 需要二刷
func constructFromPrePost(preorder []int, postorder []int) (ans *TreeNode) {
	var getTree func(preStart int, preEnd int, postStart int, postEnd int) *TreeNode
	getTree = func(preStart int, preEnd int, postStart int, postEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}

		currentLength := preEnd - preStart + 1
		root := &TreeNode{Val: preorder[preStart]}
		if currentLength == 1 {
			return root
		}

		breakPointIndex := postStart
		for breakPointIndex <= postEnd && postorder[breakPointIndex] != preorder[preStart+1] {
			breakPointIndex += 1
		}

		leftCount := breakPointIndex - postStart + 1
		root.Left = getTree(preStart+1, preStart+leftCount, postStart, postStart+breakPointIndex-1)
		root.Right = getTree(preStart+leftCount+1, preEnd, postStart+leftCount, postEnd-1)

		return root
	}
	return getTree(0, len(preorder)-1, 0, len(postorder)-1)
}

// @lc code=end
//      1
//   2     3
//  4 5   6 7
// pre  :                       1 [ *2* 4 5] [3 6 7]
// post :                        [4 5 *2*] [6 7 3] 1
// breakPointIndex  post :             ^
//  leftCount = 2 - 0 + 1 = 3
//
//

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// preorder := Deserialize[[]int](ReadLine(stdin))
	// postorder := Deserialize[[]int](ReadLine(stdin))

	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	ans := constructFromPrePost(preorder, postorder)

	fmt.Println("\noutput:", Serialize(ans))
}
