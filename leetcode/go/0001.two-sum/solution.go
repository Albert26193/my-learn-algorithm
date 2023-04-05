// Created by Albert at 2023/04/04 23:04
// https://leetcode.cn/problems/two-sum/

package main

import (
	//"bufio"
	"fmt"
	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	return
}

// @lc code=end

func main() {
	//stdin := bufio.NewReader(os.Stdin)
	// nums := Deserialize[[]int](ReadLine(stdin))
	nums := []int{1, 2, 3, 4, 5}
	// target := Deserialize[int](ReadLine(stdin))
	target := 200
	ans := twoSum(nums, target)
	fmt.Println("output: " + Serialize(ans))
}
