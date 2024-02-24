package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   v,
	}
}

func main() {
	n1 := NewTreeNode(3)
	n2 := NewTreeNode(9)
	n3 := NewTreeNode(20)
	n4 := NewTreeNode(15)
	n5 := NewTreeNode(7)

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	ans := levelOrder(n1)
	fmt.Println(ans)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := make([]*TreeNode, 0)

	ans := make([][]int, 0)

	q = append(q, root)

	cnt := 0
	for len(q) > 0 {
		levelSize := len(q)
		ans = append(ans, []int{})

		for i := 0; i < levelSize; i++ {
			head := q[0]
			q = q[1:]
			ans[cnt] = append(ans[cnt], head.Val)
			// fmt.Println(head)

			if head.Left != nil {
				q = append(q, head.Left)
			}

			if head.Right != nil {
				q = append(q, head.Right)
			}
		}
		cnt += 1
	}

	// fmt.Println(ans, cnt)
	return ans
}
