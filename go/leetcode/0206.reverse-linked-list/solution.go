// Created by Albert's server at 2024/02/20 18:33
// leetgo: dev
// https://leetcode.cn/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return head
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
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
