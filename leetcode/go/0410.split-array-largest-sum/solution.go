// Created by Albert's server at 2024/01/21 21:56
// leetgo: dev
// https://leetcode.cn/problems/split-array-largest-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func splitArray(nums []int, k int) (ans int) {
	n := len(nums)
	sum := make([]int, n+1)
	maxNum := nums[0]
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
		if nums[i-1] > maxNum {
			maxNum = nums[i-1]
		}
	}
	var check = func(cap int) bool {
		prev, cnt := 0, 1
		for i := 0; i <= n; i++ {
			if sum[i]-sum[prev] > cap {
				prev = i - 1
				cnt += 1
			}
		}

		return cnt <= k
	}

	left, right := maxNum, sum[n]
	for left < right {
		mid := (right + left) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
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
	ans := splitArray(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
