// Created by Albert's server at 2024/01/13 23:36
// leetgo: dev
// https://leetcode.cn/problems/minimum-cost-for-tickets/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mincostTickets(days []int, costs []int) (ans int) {
	f := make([]int, 366)

	sort.Ints(days)
	n := len(days)

	for i := 1; i < 366; i++ {
		f[i] = 0x3f3f3f3f
	}

	f[0] = 0
	f[1] = minn(costs...)
	for i, index := 1, 0; index <= n-1 && i < 366; i++ {
		if days[index] == i {
			index += 1
			f[i] = minn(f[maxx(0, i-1)]+costs[0], f[maxx(0, i-7)]+costs[1], f[maxx(0, i-30)]+costs[2])
			ans = f[i]
		} else {
			f[i] = f[i-1]
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
	days := Deserialize[[]int](ReadLine(in))
	costs := Deserialize[[]int](ReadLine(in))
	ans := mincostTickets(days, costs)

	fmt.Println("\noutput:", Serialize(ans))
}
