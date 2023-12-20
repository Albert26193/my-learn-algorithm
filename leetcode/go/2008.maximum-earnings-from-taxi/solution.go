// Created by Albert's server at 2023/12/20 13:26
// leetgo: dev
// https://leetcode.cn/problems/maximum-earnings-from-taxi/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// f[i] = m, end with i, max profit is m
// f[i] = f[i-1]
// f[i] = f[k] + profit[k] ==> k can get max profit
func maxTaxiEarnings(n int, rides [][]int) (ans int64) {
	rec := make(map[int][][]int)

	for _, ride := range rides {
		rec[ride[1]] = append(rec[ride[1]], ride)
	}

	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		for _, ride := range rec[i] {
			f[i] = maxx(f[i], f[ride[0]]+ride[2]+ride[1]-ride[0])
		}
	}

	return int64(maxx(f...))
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
	n := Deserialize[int](ReadLine(in))
	rides := Deserialize[[][]int](ReadLine(in))
	ans := maxTaxiEarnings(n, rides)

	fmt.Println("\noutput:", Serialize(ans))
}
