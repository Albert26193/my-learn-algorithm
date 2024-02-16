// Created by Albert's server at 2024/02/16 11:22
// leetgo: dev
// https://leetcode.cn/problems/longest-increasing-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// greedy + binary search
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	cur := make([]int, 0)
	cur = append(cur, nums[0])

	for _, v := range nums[1:] {
		if v > cur[len(cur)-1] {
			cur = append(cur, v)
			continue
		}

		l, r := 0, len(cur)-1

		for l < r {
			mid := (l + r) / 2
			if v <= cur[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if l < len(cur) && v <= cur[l] {
			cur[l] = v
		}
	}

	return len(cur)
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
	ans := lengthOfLIS(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
