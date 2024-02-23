// Created by Albert's server at 2024/02/22 17:15
// leetgo: dev
// https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSubArrayLen(nums []int, k int) (ans int) {
	n := len(nums)

	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// k: preSum[i]
	// v: index
	mp := make(map[int]int)
	// mp[0] = -1
	for i := 0; i <= n; i++ {
		if preSum[i] == k {
			ans = i
		}

		if _, ok := mp[preSum[i]-k]; ok {
			ans = maxx(ans, i-mp[preSum[i]-k])
		}

		if _, ok := mp[preSum[i]]; !ok {
			mp[preSum[i]] = i
		}
	}

	return
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
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
	ans := maxSubArrayLen(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
