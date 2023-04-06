// Created by Albert at 2023/04/05 21:09
// https://leetcode.cn/problems/number-of-common-factors/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func commonFactors(a int, b int) (ans int) {
	i := 1
	for i <= 1005 {
		if (a%i == 0) && (b%i == 0) {
			ans += 1
		}

		i += 1
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	a := Deserialize[int](ReadLine(stdin))
	b := Deserialize[int](ReadLine(stdin))
	ans := commonFactors(a, b)
	fmt.Println("output: " + Serialize(ans))
}
