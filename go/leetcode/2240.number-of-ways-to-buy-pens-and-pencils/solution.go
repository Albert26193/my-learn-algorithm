// Created by Albert at 2023/09/01 19:01
// leetgo: 1.3.7
// https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func waysToBuyPensPencils(total int, cost1 int, cost2 int) (ans int64) {
	// 确保 cost1 比较大,枚举起来快
	if cost1 < cost2 {
		return waysToBuyPensPencils(total, cost2, cost1)
	}

	res := 0
	cnt := 0
	for cnt*cost1 <= total {
		// 因为一个不选也是一种方案
		// 比如剩余 5 可以让 cost2(val = 3) 选择
		// 5 // 3 = 1 ... 1
		// 2 种方案: 0个 1个
		res += (total-cnt*cost1)/cost2 + 1
		cnt += 1
	}
	return int64(res)
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// total := Deserialize[int](ReadLine(stdin))
	// cost1 := Deserialize[int](ReadLine(stdin))
	// cost2 := Deserialize[int](ReadLine(stdin))
	total := 20
	cost1 := 10
	cost2 := 5
	ans := waysToBuyPensPencils(total, cost1, cost2)

	fmt.Println("\noutput:", Serialize(ans))
}
