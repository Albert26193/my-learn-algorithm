// Created by Albert's server at 2023/12/08 15:05
// leetgo: dev
// https://leetcode.cn/problems/maximum-earnings-from-taxi/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// f[i][j] = p
// i: from [0, ... , i-1]
// j: pick j person
// f[i][0] ==> not have a passenger
// f[i][1] ==> have a passenger
// forbidden

// get in ==> pay
func maxTaxiEarnings(n int, rides [][]int) (ans int64) {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})

	fmt.Println(rides)

	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, 2)
	}

	maxReach := 0
	for i := 1; i <= n; i++ {
		start := rides[i-1][0]
		end := rides[i-1][1]
		currentProfit := end - start + rides[i-1][2]

		if start >= maxReach {
			f[i][0] = maxx(f[i-1][0], f[i-1][1])
			f[i][1] = maxx(f[i-1][1], f[i-1][0]+currentProfit)
		} else {
			f[i][0] = f[i-1][0]
			f[i][1] = f[i-1][1]
		}
	}

	return int64(maxx(f[n]...))
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
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
	n := Deserialize[int](ReadLine(in))
	rides := Deserialize[[][]int](ReadLine(in))
	fmt.Println(n, rides)
	ans := maxTaxiEarnings(n, rides)

	fmt.Println("\noutput:", Serialize(ans))
}
