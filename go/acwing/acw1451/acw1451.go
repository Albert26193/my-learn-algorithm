package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1, n2, n3 := &ListNode{Val: 5}, &ListNode{Val: 7}, &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n3
	printList(n1)
	n1 = quickSort(n1)
	printList(n1)
}

func getTail(node *ListNode) *ListNode {
	for node != nil && node.Next != nil {
		node = node.Next
	}

	return node
}

// core function
func quickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pivot := head.Val

	lowDummy, midDummy, highDummy := &ListNode{Val: -1}, &ListNode{Val: -1}, &ListNode{Val: -1}
	low, mid, high := lowDummy, midDummy, highDummy

	p := head
	for p != nil {
		if p.Val < pivot {
			low.Next = p
			low = low.Next
		} else if p.Val == pivot {
			mid.Next = p
			mid = mid.Next
		} else {
			high.Next = p
			high = high.Next
		}

		p = p.Next
	}

	// 核心： 最后需要切断链表
	low.Next, mid.Next, high.Next = nil, nil, nil

	lowDummy.Next = quickSort(lowDummy.Next)
	highDummy.Next = quickSort(highDummy.Next)

	getTail(lowDummy).Next = midDummy.Next
	getTail(midDummy).Next = highDummy.Next

	return lowDummy.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
