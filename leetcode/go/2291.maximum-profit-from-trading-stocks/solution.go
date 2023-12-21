// Created by Albert's server at 2023/12/21 20:28
// leetgo: dev
// https://leetcode.cn/problems/maximum-profit-from-trading-stocks/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumProfit(present []int, future []int, budget int) (ans int) {
	fmt.Println(present, future, budget)
	n := len(present)

	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, budget+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= budget; j++ {
			if j < present[i-1] {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = maxx(f[i-1][j], f[i-1][j-present[i-1]]+future[i-1]-present[i-1])
			}
		}
	}

	return f[n][budget]
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
	present := Deserialize[[]int](ReadLine(in))
	future := Deserialize[[]int](ReadLine(in))
	budget := Deserialize[int](ReadLine(in))
	ans := maximumProfit(present, future, budget)

	fmt.Println("\noutput:", Serialize(ans))
}
