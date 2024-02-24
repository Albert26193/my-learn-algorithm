// Created by Albert's server at 2024/01/06 19:04
// leetgo: dev
// https://leetcode.cn/problems/path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasPathSum(root *TreeNode, targetSum int) bool {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	targetSum := Deserialize[int](ReadLine(stdin))
	ans := hasPathSum(root, targetSum)

	fmt.Println("\noutput:", Serialize(ans))
}
