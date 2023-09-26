// Created by Albert at 2023/09/26 14:43
// leetgo: 1.3.8
// https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/
// https://leetcode.cn/contest/weekly-contest-355/problems/largest-element-in-an-array-after-merge-operations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxArrayValue(nums []int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxArrayValue(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
