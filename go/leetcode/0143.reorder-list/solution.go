// Created by Albert's server at 2024/02/20 16:50
// leetgo: dev
// https://leetcode.cn/problems/reorder-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 1. find middle node
// 2. split and reverse
// 3. join List
func findMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pa, pb := head, head

	for pb.Next != nil && pb.Next.Next != nil {
		pa = pa.Next
		pb = pb.Next.Next
	}

	return pa
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var cur = head
	var prev *ListNode

	for cur != nil {
		nextNode := cur.Next
		cur.Next = prev

		prev = cur
		cur = nextNode
	}

	return prev
}

func joinList(ha *ListNode, hb *ListNode) {
	pa, pb := ha, hb

	for pa != nil && pb != nil {
		tempA, tempB := pa.Next, pb.Next
		pa.Next = pb
		pb.Next = tempA

		pa, pb = tempA, tempB
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := findMiddleNode(head)
	l1, l2 := head, mid.Next

	mid.Next = nil
	l2 = reverseList(l2)
	joinList(l1, l2)
	fmt.Println(l1.Val, l2.Val)
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	head := Deserialize[*ListNode](ReadLine(in))
	reorderList(head)
	ans := head

	fmt.Println("\noutput:", Serialize(ans))
}
