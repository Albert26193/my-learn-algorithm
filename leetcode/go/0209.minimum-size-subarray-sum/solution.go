// Created by Albert's server at 2024/02/22 15:28
// leetgo: dev
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	l, r := 0, 0

	n := len(nums)
	curSum := 0
	ans = math.MaxInt32
	for r = 0; r < n; r++ {
		curSum += nums[r]

		for l <= r && curSum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}
			curSum -= nums[l]
			l++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
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

	target := Deserialize[int](ReadLine(in))
	nums := Deserialize[[]int](ReadLine(in))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
