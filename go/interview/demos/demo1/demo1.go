package main

import "fmt"

type TreeNode struct {
	left, right *TreeNode
	val         int
}

func NewTreeNode(n int) *TreeNode {
	return &TreeNode{
		left:  nil,
		right: nil,
		val:   n,
	}
}

func main() {
	var n1 = NewTreeNode(1)
	var n2 = NewTreeNode(2)
	var n3 = NewTreeNode(3)
	var n4 = NewTreeNode(4)
	var n5 = NewTreeNode(5)
	n1.left = n2
	n1.right = n5
	n2.left = n3
	n3.right = n4

	fmt.Println(n1.left.left.val)
}
