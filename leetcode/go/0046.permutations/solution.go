// Created by Albert at 2023/04/19 00:29
// https://leetcode.cn/problems/permutations/

package main

import (
	"fmt"
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func backTrack(nums []int, curIndex int, combination []int, ans *[][]int, vis []int) {
	if curIndex > len(nums) {
		return
	}

	if curIndex == len(nums) {
		*ans = append(*ans, combination)
		return
	}

	for i := range nums {
		if vis[i] == 1 {
			continue
		}

		vis[i] = 1
		combination = append(combination, i+1)
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
	//stdin := bufio.NewReader(os.Stdin)
	//nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{1, 2, 3}
	ans := permute(nums)
	fmt.Println("output: " + Serialize(ans))
}
