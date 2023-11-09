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
	lenNums := len(nums)
	preSum := make([]int, lenNums+1)

	for i := 1; i <= lenNums; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// record: key: currentSum % k
	// record: value: count for key
	record := make(map[int]int)
	for i := 0; i <= lenNums; i++ {
		currentSum := preSum[i]
		rest := (currentSum%k + k) % k
		if _, exist := record[rest]; exist {
			ans += record[rest]
		}

		record[rest] += 1
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
