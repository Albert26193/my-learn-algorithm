// Created by Albert's server at 2024/01/03 22:46
// leetgo: dev
// https://leetcode.cn/problems/rotate-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = (n + k) % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
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
	k := Deserialize[int](ReadLine(in))
	rotate(nums, k)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
