// Created by Albert at 2023/11/09 11:41
// leetgo: 1.3.8
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func subarraysDivByK(nums []int, k int) (ans int) {
	n := len(nums)

	preSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// k: preSum % k
	// v: cnt
	mp := make(map[int]int)
	rest := 0
	for i := 0; i <= n; i++ {
		rest = (k + preSum[i]%k) % k
		if _, ok := mp[rest]; ok {
			ans += mp[rest]
		}

		mp[rest] += 1
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
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	nums := Deserialize[[]int](ReadLine(reader))
	k := Deserialize[int](ReadLine(reader))
	ans := subarraysDivByK(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
