// Created by Albert at 2023/11/09 11:10
// leetgo: 1.3.8
// https://leetcode.cn/problems/subarray-sum-equals-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func subarraySum(nums []int, k int) int {
	n := len(nums)

	preSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// mp: k: prefix sum  v: count
	mp := make(map[int]int)
	cnt := 0
	for i := 0; i <= n; i++ {
		if _, ok := mp[preSum[i]-k]; ok {
			cnt += mp[preSum[i]-k]
		}
		mp[preSum[i]] += 1
	}

	return cnt
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	nums := Deserialize[[]int](ReadLine(reader))
	k := Deserialize[int](ReadLine(reader))
	ans := subarraySum(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
