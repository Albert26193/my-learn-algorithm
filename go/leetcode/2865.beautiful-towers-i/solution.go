// Created by Albert's server at 2024/01/24 09:18
// leetgo: dev
// https://leetcode.cn/problems/beautiful-towers-i/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSumOfHeights(maxHeights []int) (ans int64) {
	n := len(maxHeights)
	for i := range maxHeights {
		pre, suf := maxHeights[i], maxHeights[i]
		psum := int64(maxHeights[i])

		for j := i - 1; j >= 0; j-- {
			pre = minn(pre, maxHeights[j])
			psum += int64(pre)
		}

		for j := i + 1; j < n; j++ {
			suf = minn(suf, maxHeights[j])
			psum += int64(suf)
		}

		ans = maxx(ans, psum)
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

func maxx(nums ...int64) int64 {
	res := nums[0]
	for _, v := range nums {
		if v > res {
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
	maxHeights := Deserialize[[]int](ReadLine(in))
	ans := maximumSumOfHeights(maxHeights)

	fmt.Println("\noutput:", Serialize(ans))
}
