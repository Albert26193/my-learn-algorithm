package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
		Next: nil,
	}
}

func main() {
	n0 := NewListNode(0)
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	n4 := NewListNode(4)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	ans := Traversal(n0)
	fmt.Println(ans)
}

func Traversal(node *ListNode) []int {
	ans := make([]int, 0)

	p := node

	for p != nil {
		ans = append(ans, p.Val)
		p = p.Next
	}
	return ans
}
