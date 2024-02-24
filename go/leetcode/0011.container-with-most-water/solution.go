// Created by Albert's server at 2024/01/03 17:00
// leetgo: dev
// https://leetcode.cn/problems/container-with-most-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		currentLength := right - left

		currentVolume := currentLength * minn(height[left], height[right])
		ans = maxx(currentVolume, ans)

		// lower ==> change
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return
}

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if res > v {
			res = v
		}

	}
	return res
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if res < v {
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
	ans := maxArea(height)

	fmt.Println("\noutput:", Serialize(ans))
}
