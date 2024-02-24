// Created by Albert's server at 2024/01/06 22:01
// leetgo: dev
// https://leetcode.cn/problems/target-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func summ(nums ...int) int {
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func findTargetSumWays(nums []int, target int) (ans int) {
	// f[i] ==> sum is i, method count is f[i]

	// a - b = target
	// a + b = sum
	// a = (target + sum) / 2
	t := summ(nums...) + target
	if t%2 != 0 || t < 0 {
		return 0
	}

	t = t / 2
	n := len(nums)
	f := make([]int, t+1)

	f[0] = 1
	for i := 0; i < n; i++ {
		cur := nums[i]
		for j := t; j >= cur; j-- {
			f[j] += f[j-cur]
		}
	}

	return f[t]
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
	target := Deserialize[int](ReadLine(in))
	ans := findTargetSumWays(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
