// Created by Albert at 2023/09/17 14:13
// leetgo: 1.3.8
// https://leetcode.cn/problems/house-robber-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// [start, end]
func robRange(start int, end int, nums []int) int {
	prev, curr := 0, 0
	for i := start; i <= end; i++ {
		prev, curr = curr, maxx(curr, prev+nums[i])
	}

	return curr
}

func rob(nums []int) (ans int) {
	// choose nums[0] or 'not choose nums[0]'
	numsLength := len(nums)
	return maxx(nums[0]+robRange(2, numsLength-2, nums), robRange(1, numsLength-1, nums))
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

	stdin := bufio.NewReader(reader)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := rob(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
