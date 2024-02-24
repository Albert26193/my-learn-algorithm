// Created by Albert at 2024/10/31 13:12
// leetgo: 1.3.8
// https://leetcode.cn/problems/partition-equal-subset-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// from [0, ... , last -1], last choose ? counts, can fit sum = 0.5 * allSum
// f[i][j] = m => [0, ..., i - 1], sum = j, have m methods

// f[i][j] = f[i - 1][j]
// or f[i][j] = f[i - 1][j - nums[i - 1]]

func getSum(nums []int) int {
	cnt := 0
	for _, num := range nums {
		cnt += num
	}

	return cnt
}

// 01 back package
func canPartition(nums []int) bool {
	n := len(nums)

	V := getSum(nums)
	f := make([]bool, V+1)

	if V%2 != 0 {
		return false
	}

	target := V / 2
	f[0] = true
	for i := 0; i < n; i++ {
		currentVolume := nums[i]
		for j := target; j >= currentVolume; j-- {
			if f[target] {
				return true
			}

			f[j] = f[j] || f[j-nums[i]]
		}
	}

	return f[target]
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
	ans := canPartition(nums)
	fmt.Println("\noutput:", Serialize(ans))
}
