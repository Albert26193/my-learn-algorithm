// Created by Albert at 2023/04/19 02:00
// https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO
func minDays(bloomDay []int, m int, k int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	bloomDay := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := minDays(bloomDay, m, k)
	fmt.Println("output: " + Serialize(ans))
}
