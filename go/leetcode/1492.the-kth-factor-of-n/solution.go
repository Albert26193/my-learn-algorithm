// Created by Albert at 2023/10/23 11:38
// leetgo: 1.3.8
// https://leetcode.cn/problems/the-kth-factor-of-n/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func kthFactor(n int, k int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := kthFactor(n, k)

	fmt.Println("\noutput:", Serialize(ans))
}
