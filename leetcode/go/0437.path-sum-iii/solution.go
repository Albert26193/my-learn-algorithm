// Created by Albert's server at 2024/01/06 15:05
// leetgo: dev
// https://leetcode.cn/problems/path-sum-iii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func pathSum(root *TreeNode, targetSum int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	targetSum := Deserialize[int](ReadLine(stdin))
	ans := pathSum(root, targetSum)

	fmt.Println("\noutput:", Serialize(ans))
}
