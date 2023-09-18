// Created by Albert at 2023/09/18 10:56
// leetgo: 1.3.8
// https://leetcode.cn/problems/happy-students/
// https://leetcode.cn/contest/weekly-contest-363/problems/happy-students/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countWays(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := countWays(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
