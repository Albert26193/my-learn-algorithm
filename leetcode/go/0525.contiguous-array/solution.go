// Created by Albert's server at 2024/02/22 16:58
// leetgo: dev
// https://leetcode.cn/problems/contiguous-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMaxLength(nums []int) (ans int) {
	// n := len(nums)

	mp := make(map[int]int)
	cnt := 0
	for i, num := range nums {
		if num == 0 {
			cnt -= 1
		} else {
			cnt += 1
		}

		if cnt == 0 {
			ans = i + 1
		}

		if _, ok := mp[cnt]; ok {
			ans = maxx(ans, i-mp[cnt])
		} else {
			mp[cnt] = i
		}
	}
	return
}

func maxx(nums ...int) int {
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
	nums := Deserialize[[]int](ReadLine(in))
	ans := findMaxLength(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
