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

// level: k
//  root
// 0:   1 			2 3 ... n
// 1:  2 	3...n   	...
// 2: 3...n ...         ...
func backTrack(n int, k int, currentIndex int, combination *[]int, ans *[][]int) {
	if len(*combination) == k {
		copyCombination := make([]int, k)
		copy(copyCombination, *combination)
		*ans = append(*ans, copyCombination)
		return
	}

	for i := currentIndex; i <= n; i++ {
		*combination = append(*combination, i)
		backTrack(n, k, i+1, combination, ans)
		*combination = (*combination)[:len(*combination)-1]
	}
}

func combine(n int, k int) (ans [][]int) {
	combination := make([]int, 0)
	backTrack(n, k, 1, &combination, &ans)
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
