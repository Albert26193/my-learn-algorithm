// Created by Albert at 2023/09/01 18:00
// leetgo: 1.3.7
// https://leetcode.cn/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {
	n := len(height)

	// left[i] = j ==> all nums left i, max is j
	left, right := make([]int, n), make([]int, n)

	leftMax, rightMax := 0, 0
	for i := 0; i < n; i++ {
		curLeft, curRight := height[i], height[n-1-i]
		if curLeft > leftMax {
			leftMax = curLeft
		}
		left[i] = leftMax
		if curRight > rightMax {
			rightMax = curRight
		}
		right[n-1-i] = rightMax
	}

	for i := 0; i < n; i++ {
		ans += (minn(left[i], right[i]) - height[i])
	}
	return
}

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
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
	height := Deserialize[[]int](ReadLine(in))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
