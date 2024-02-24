// Created by Albert at 2023/09/06 16:55
// leetgo: 1.3.7
// https://leetcode.cn/problems/serialize-and-deserialize-bst/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Codec struct {
    
}

func Constructor() (ans Codec) {

	return
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) (ans *TreeNode) {

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := CodecDriver(root)

	fmt.Println("\noutput:", Serialize(ans))
}
