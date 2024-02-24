// Created by Albert's server at 2024/02/21 14:20
// leetgo: dev
// https://leetcode.cn/problems/wiggle-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func wiggleMaxLength(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	nums := Deserialize[[]int](ReadLine(in))
	ans := wiggleMaxLength(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
