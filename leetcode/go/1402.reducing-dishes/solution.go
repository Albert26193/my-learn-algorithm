// Created by Albert at 2023/10/23 11:38
// leetgo: 1.3.8
// https://leetcode.cn/problems/reducing-dishes/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func maxx(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]

	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func maxSatisfaction(satisfaction []int) (ans int) {
	sort.Ints(satisfaction)
	// fmt.Println(satisfaction)

	cooksLen := len(satisfaction)
	dp := make([][]int, cooksLen+1)

	for i := range dp {
		dp[i] = make([]int, cooksLen+1)
	}

	// dp[i][j]: from 0...i-1(count: i) choose j => maxSatisfaction
	/*
	   'j'th pick ,if not choose index 'i':
	       dp[i][j] = max(dp[i - 1][j], dp[i][j])
	   if choose it:
	       dp[i][j] = max(dp[i][j], dp[i - 1][j - 1] + j * satisfaction[i - 1])
	*/

	dp[0][0] = 0
	dp[1][0] = 0

	for i := 1; i <= cooksLen; i++ {
		for j := 1; j <= i; j++ {
			if j <= i-1 {
				dp[i][j] = maxx(dp[i-1][j], dp[i-1][j-1]+j*satisfaction[i-1])
			} else {
				dp[i][j] = dp[i-1][j-1] + j*satisfaction[i-1]
			}
		}
	}

	return maxx(dp[cooksLen]...)
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

	satisfaction := Deserialize[[]int](ReadLine(reader))
	ans := maxSatisfaction(satisfaction)

	fmt.Println("\noutput:", Serialize(ans))
}
