// Created by Albert's server at 2024/02/18 14:02
// leetgo: dev
// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func preorder(root *Node) (ans []int) {

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[int](ReadLine(stdin))
	ans := preorder(root)

	fmt.Println("\noutput:", Serialize(ans))
}
