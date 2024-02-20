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

func reorderList(head *ListNode)  {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	reorderList(head)
	ans := head

	fmt.Println("\noutput:", Serialize(ans))
}
