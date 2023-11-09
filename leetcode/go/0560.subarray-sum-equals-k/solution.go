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
	lenNUms := len(nums)
	preSum := make([]int, lenNUms+1)

	for i := 1; i <= lenNUms; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// key: sum
	// value: count
	record := make(map[int]int)
	cnt := 0
	for i := 0; i <= lenNUms; i++ {
		currentSum := preSum[i]
		if _, exist := record[currentSum-k]; exist {
			cnt += record[currentSum-k]
		}

		record[preSum[i]] += 1
	}

	// fmt.Println(preSum)
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
