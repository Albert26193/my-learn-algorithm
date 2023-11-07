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

func canPartition(nums []int) bool {
	numsLength := len(nums)
	totalSum := getSum(nums)

	if totalSum%2 == 1 {
		return false
	}

	if numsLength == 0 || numsLength == 1 {
		return false
	}

	targetSum := totalSum / 2

	f := make([][]bool, numsLength+1)
	for i := 0; i <= numsLength; i++ {
		f[i] = make([]bool, totalSum+1)
	}

	// boundary:
	f[0][0] = true
	for i := 1; i <= numsLength; i++ {
		f[i][0] = true
		f[i][nums[i-1]] = true
	}

	for i := 1; i <= numsLength; i++ {
		currentNum := nums[i-1]
		for j := 0; j <= targetSum; j++ {
			if j >= currentNum {
				f[i][j] = f[i-1][j] || f[i-1][j-currentNum]
			} else {
				f[i][j] = f[i-1][j]
			}
		}

		if f[i][targetSum] {
			return true
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
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)
	nums := Deserialize[[]int](ReadLine(reader))
	ans := canPartition(nums)
	fmt.Println("\noutput:", Serialize(ans))
}
