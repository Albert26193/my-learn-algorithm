// Created by Albert's server at 2024/01/24 10:00
// leetgo: dev
// https://leetcode.cn/problems/beautiful-towers-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSumOfHeights(maxHeights []int) (ans int64) {
	sl, sr := make([]int, 0), make([]int, 0)
	n := len(maxHeights)
	pre, suf := make([]int64, n), make([]int64, n)

	for i := 0; i < n; i++ {
		for len(sl) > 0 && maxHeights[sl[len(sl)-1]] > maxHeights[i] {
			sl = sl[:len(sl)-1]
		}

		if len(sl) == 0 {
			pre[i] = int64((i + 1) * maxHeights[i])
		} else {
			top := sl[len(sl)-1]
			pre[i] = pre[top] + int64(maxHeights[i]*(i-top))
		}

		sl = append(sl, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(sr) > 0 && maxHeights[sr[len(sr)-1]] > maxHeights[i] {
			sr = sr[:len(sr)-1]
		}

		if len(sr) == 0 {
			suf[i] = int64((n - i) * maxHeights[i])
		} else {
			top := sr[len(sr)-1]
			suf[i] = suf[top] + int64(maxHeights[i]*(top-i))
		}

		sr = append(sr, i)
		ans = maxx(ans, pre[i]+suf[i]-int64(maxHeights[i]))
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
