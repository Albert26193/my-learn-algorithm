// Created by Albert's server at 2024/02/22 16:41
// leetgo: dev
// https://leetcode.cn/problems/subarray-product-less-than-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	n := len(nums)

	l, r := 0, 0
	product := 1
	for r = 0; r < n; r++ {
		product *= nums[r]

		for l <= r && product >= k {
			product = product / nums[l]
			l++
		}

		ans += (r - l + 1)
	}
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
	k := Deserialize[int](ReadLine(in))
	ans := numSubarrayProductLessThanK(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
