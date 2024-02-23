// Created by Albert at 2023/07/08 15:25
// leetgo: 1.3.2
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func constructFromPrePost(preorder []int, postorder []int) (ans *TreeNode) {
	var getTree func(pre []int, post []int) *TreeNode
	getTree = func(pre, post []int) *TreeNode {
		if len(pre) == 0 || len(post) == 0 || len(pre) != len(post) {
			return nil
		}

		n := len(pre)
		root := &TreeNode{
			Val: pre[0],
		}

		if n == 1 {
			return root
		}

		leftLen := 0
		for i := 0; i < n; i++ {
			if post[i] == pre[1] {
				leftLen = i + 1
				break
			}
		}

		root.Left = getTree(pre[1:leftLen+1], post[:leftLen])
		root.Right = getTree(pre[leftLen+1:], post[leftLen:len(post)-1])

		return root
	}

	ans = getTree(preorder, postorder)
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()
	ReadLine(in)
	preorder := Deserialize[[]int](ReadLine(in))
	postorder := Deserialize[[]int](ReadLine(in))
	ans := constructFromPrePost(preorder, postorder)

	fmt.Println("\noutput:", Serialize(ans))
}
