// Created by Albert at 2023/11/12 23:21
// leetgo: 1.3.8
// https://leetcode.cn/problems/next-greater-element-i/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func nextGreaterElement(nums1 []int, nums2 []int) (ans []int) {
	mp := make(map[int]int)

	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		// 1. pop
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// 2. stack top is the element to find
		if len(stack) == 0 {
			mp[nums2[i]] = -1
		} else {
			mp[nums2[i]] = nums2[stack[len(stack)-1]]
		}

		// 3. push
		stack = append(stack, i)
	}

	for i := 0; i < len(nums1); i++ {
		ans = append(ans, mp[nums1[i]])
	}

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

	nums1 := Deserialize[[]int](ReadLine(reader))
	nums2 := Deserialize[[]int](ReadLine(reader))
	ans := nextGreaterElement(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
