// Created by Albert at 2023/09/18 10:56
// leetgo: 1.3.8
// https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/
// https://leetcode.cn/contest/weekly-contest-363/problems/maximum-element-sum-of-a-complete-subset-of-indices/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSum(nums []int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maximumSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
