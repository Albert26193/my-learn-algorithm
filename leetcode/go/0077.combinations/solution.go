// Created by Albert at 2023/04/07 16:22
// https://leetcode.cn/problems/combinations/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func backTrack(n int, k int, depth int, combination *[]int, ans *[][]int) {
	if depth == (k - 1) {
		*ans = append(*ans, *combination)
		return
	}

	for i := depth + 1; i <= n; i++ {
		*combination = append(*combination, i)
		backTrack(n, k, depth+1, combination, ans)
		// *combination = (*combination)[:len(*combination)-1]
	}
}

func combine(n int, k int) (ans [][]int) {
	backTrack(n, k, 0, &([]int{}), &ans)
	return ans
}

// @lc code=end

func main() {
	/* stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin)) */

	n := 4
	k := 2
	ans := combine(n, k)
	fmt.Println("output: " + Serialize(ans))
}
