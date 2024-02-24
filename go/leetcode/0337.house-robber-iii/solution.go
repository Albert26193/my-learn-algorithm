// Created by Albert at 2023/09/18 09:41
// leetgo: 1.3.8
// https://leetcode.cn/problems/house-robber-iii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// choose or 'not choose'
func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lRob, lNotRob := dfs(root.Left)
	rRob, rNotRob := dfs(root.Right)
	// choose current node
	rob := lNotRob + rNotRob + root.Val
	// not choose
	notRob := maxx(lRob, lNotRob) + maxx(rRob, rNotRob)
	return rob, notRob
}

func rob(root *TreeNode) (ans int) {
	return maxx(dfs(root))
}

// @lc code=end

func main() {

	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	stdin := bufio.NewReader(reader)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := rob(root)

	fmt.Println("\noutput:", Serialize(ans))
}
