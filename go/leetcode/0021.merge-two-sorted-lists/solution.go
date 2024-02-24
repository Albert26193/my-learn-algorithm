// Created by Albert's server at 2024/02/20 17:32
// leetgo: dev
// https://leetcode.cn/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var dummy = &ListNode{
		Val: -1,
	}
	var prev = dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return dummy.Next
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
	list1 := Deserialize[*ListNode](ReadLine(in))
	list2 := Deserialize[*ListNode](ReadLine(in))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}
