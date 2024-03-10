// Created by Albert's server at 2024/03/07 14:59
// leetgo: dev
// https://leetcode.cn/problems/odd-even-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func oddEvenList(head *ListNode) (ans *ListNode) {
	p := head

	evenHead, oddHead := &ListNode{}, &ListNode{}
	even, odd := evenHead, oddHead

	isOdd := true
	for p != nil {
		if !isOdd {
			even.Next = p
			even = even.Next
		} else {
			odd.Next = p
			odd = odd.Next
		}

		p = p.Next
		isOdd = !isOdd
	}

	odd.Next = evenHead.Next
	even.Next = nil

	return oddHead.Next
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
	ans := oddEvenList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
