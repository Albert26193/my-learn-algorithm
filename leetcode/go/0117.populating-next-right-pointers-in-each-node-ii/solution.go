// Created by Albert at 2023/11/11 00:29
// leetgo: 1.3.8
// https://leetcode.cn/problems/populating-next-Right-pointers-in-each-node-ii/

package main

type Node struct {
	val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// @lc code=begin

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			if i == levelSize-1 {
				cur.Next = nil
			} else {
				cur.Next = queue[0]
			}

		}
	}

	return root
}

// @lc code=end
