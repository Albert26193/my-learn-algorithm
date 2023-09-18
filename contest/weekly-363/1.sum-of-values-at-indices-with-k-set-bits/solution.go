// Created by Albert at 2023/09/18 10:56
// leetgo: 1.3.8
// https://leetcode.cn/problems/sum-of-values-at-indices-with-k-set-bits/
// https://leetcode.cn/contest/weekly-contest-363/problems/sum-of-values-at-indices-with-k-set-bits/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sumIndicesWithKSetBits(nums []int, k int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := sumIndicesWithKSetBits(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
