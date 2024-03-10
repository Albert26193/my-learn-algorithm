// Created by Albert's server at 2024/03/07 13:27
// leetgo: dev
// https://leetcode.cn/problems/add-two-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	p1, p2 := l1, l2

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	carry := 0
	var header, pointer *ListNode

	pointer = header
	for p1 != nil || p2 != nil {
		n1, n2 := 0, 0

		if p1 != nil {
			n1 = p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			n2 = p2.Val
			p2 = p2.Next
		}

		curSum := n1 + n2 + carry
		sum := curSum % 10
		carry = curSum / 10

		if header == nil {
			header = &ListNode{
				Val: sum,
			}
			pointer = header
		} else {
			pointer.Next = &ListNode{
				Val: sum,
			}
			pointer = pointer.Next
		}
	}

	if carry > 0 {
		pointer.Next = &ListNode{
			Val: carry,
		}
	}

	ans = header
	return
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
	l1 := Deserialize[*ListNode](ReadLine(in))
	l2 := Deserialize[*ListNode](ReadLine(in))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
