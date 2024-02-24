// Created by Albert's server at 2024/02/20 18:28
// leetgo: dev
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// head , .... xxx, tail
// xxx, ....,   head
// TODO: fix it
func reverseKGroup(head *ListNode, k int) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	revHead := reverseList(head, head.Next)
	fmt.Println(revHead)
	printList(revHead)

	return
}

// return new head
func reverseList(head *ListNode, tail *ListNode) *ListNode {
	var cur = head
	var prev *ListNode

	for cur.Next != tail {
		nextNode := cur.Next
		cur.Next = prev

		prev = cur
		cur = nextNode
	}

	return cur
}

func printList(head *ListNode) {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	fmt.Println(list)
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
	k := Deserialize[int](ReadLine(in))
	ans := reverseKGroup(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}
