// Created by Albert at 2023/04/13 19:57
// https://leetcode.cn/problems/subsets/

package main

import (
	// "bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func backTrack(nums []int, ans *[][]int, startIndex int, combination []int) {
	if startIndex > len(nums) {
		return
	}

	// 注意这里需要复制 combination
	temp := make([]int, len(combination))
	copy(temp, combination)
	*ans = append(*ans, temp)

	for i := startIndex; i < len(nums); i++ {
		combination = append(combination, nums[i])
		backTrack(nums, ans, i+1, combination)
		combination = combination[:len(combination)-1]
	}
}

func subsets(nums []int) (ans [][]int) {
	backTrack(nums, &ans, 0, []int{})
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{1, 2, 3}
	ans := subsets(nums)
	fmt.Println("output: " + Serialize(ans))
}
