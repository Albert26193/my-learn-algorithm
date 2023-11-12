// Created by Albert at 2023/04/19 00:29
// https://leetcode.cn/problems/permutations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// redo it on 2023-11-13
func backTrack(nums []int, curIndex int, combination []int, ans *[][]int, vis []int) {
	if curIndex > len(nums) {
		return
	}

	if curIndex == len(nums) {
		copyCombination := make([]int, len(nums))
		copy(copyCombination, combination)
		*ans = append(*ans, copyCombination)
		return
	}

	for i := 0; i < len(nums); i++ {
		if vis[i] == 1 {
			continue
		}

		combination = append(combination, nums[i])
		vis[i] = 1

		backTrack(nums, curIndex+1, combination, ans, vis)

		combination = combination[:len(combination)-1]
		vis[i] = 0
	}
}

func permute(nums []int) (ans [][]int) {
	vis := make([]int, len(nums))
	backTrack(nums, 0, []int{}, &ans, vis)
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

	ans := permute(nums)
	fmt.Println("output: " + Serialize(ans))
}
