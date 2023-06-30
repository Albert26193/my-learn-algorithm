// Created by Albert at 2023/06/29 13:55
// leetgo: 1.3.2
// https://leetcode.cn/problems/maximum-subarray-min-product/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: monotonic stack + prefix sum + multiplication principle
func maxSumMinProduct(nums []int) (ans int) {
	const mod int = 1e9 + 7

	ans = (ans%mod + mod) % mod
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSumMinProduct(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
