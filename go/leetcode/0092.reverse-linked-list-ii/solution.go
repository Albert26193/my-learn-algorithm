// Created by Albert's server at 2024/01/06 21:10
// leetgo: dev
// https://leetcode.cn/problems/reverse-linked-list-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseBetween(head *ListNode, left int, right int) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	left := Deserialize[int](ReadLine(stdin))
	right := Deserialize[int](ReadLine(stdin))
	ans := reverseBetween(head, left, right)

	fmt.Println("\noutput:", Serialize(ans))
}
