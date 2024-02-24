// Created by Albert's server at 2024/02/22 15:50
// leetgo: dev
// https://leetcode.cn/problems/continuous-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)

	// k: rest
	// v: firstPosition
	mp := make(map[int]int)

	mp[0] = -1
	rest := 0
	for i := 0; i < n; i++ {
		rest = (rest + nums[i]) % k
		if _, ok := mp[rest]; ok {
			prevIndex := mp[rest]
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			mp[rest] = i
		}
	}

	return false
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
	ans := checkSubarraySum(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
