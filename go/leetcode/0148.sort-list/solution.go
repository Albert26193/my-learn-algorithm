// Created by Albert's server at 2024/02/21 14:36
// leetgo: dev
// https://leetcode.cn/problems/sort-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func sortList(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	// find middle node
	mid := findMiddle(head)
	midNext := mid.Next
	mid.Next = nil

	// fmt.Println(mid.Val, midNext.Val)
	// sort
	left := sortList(head)
	right := sortList(midNext)

	return mergeList(left, right)
}

func findMiddle(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
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
	ans := sortList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
