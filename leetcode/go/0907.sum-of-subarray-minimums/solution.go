// Created by Albert at 2023/06/29 12:53
// leetgo: 1.3.2
// https://leetcode.cn/problems/sum-of-subarray-minimums/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: monotonic stack + multiplication principle
func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7

	ans = (ans%mod + mod) % mod
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	ans := sumSubarrayMins(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
