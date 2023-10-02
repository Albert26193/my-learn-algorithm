// Created by Albert at 2023/09/30 21:59
// leetgo: 1.3.8
// https://leetcode.cn/problems/russian-doll-envelopes/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: add binary search
// some check points can not pass
func maxx(a ...int) int {
	res := a[0]
	for _, num := range a {
		if num > res {
			res = num
		}
	}

	return res
}

func maxEnvelopes(envelopes [][]int) (ans int) {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}

		return false
	})

	/*
	   dp[i] ==> max count of envelopes, end with ans[i]
	*/

	envelopesCount := len(envelopes)

	dp := make([]int, envelopesCount)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < envelopesCount; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = maxx(dp[j]+1, dp[i])
			}
		}
	}

	return maxx(dp...)
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

	envelopes := Deserialize[[][]int](ReadLine(reader))
	ans := maxEnvelopes(envelopes)

	fmt.Println("\noutput:", Serialize(ans))
}
