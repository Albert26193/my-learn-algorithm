// Created by Albert at 2023/09/18 14:58
// leetgo: 1.3.8
// https://leetcode.cn/problems/uncrossed-lines/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxUncrossedLines(nums1 []int, nums2 []int) (ans int) {
	length1 := len(nums1)
	length2 := len(nums2)

	dp := make([][]int, length1)
	for i := 0; i < length1; i++ {
		dp[i] = make([]int, length2)
	}

	// dp[i][j]   num1[i] 和 nums2[j] 最多可以产生的连线

	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}

	for i := 0; i+1 < length2; i++ {
		if nums2[i+1] == nums1[0] {
			dp[0][i+1] = 1
		} else {
			dp[0][i+1] = dp[0][i]
		}
	}

	for i := 0; i+1 < length1; i++ {
		if nums1[i+1] == nums2[0] {
			dp[i+1][0] = 1
		} else {
			dp[i+1][0] = dp[i][0]
		}
	}

	for i := 0; i+1 < length1; i++ {
		for j := 0; j+1 < length2; j++ {
			if nums1[i+1] == nums2[j+1] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxx(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[length1-1][length2-1]
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
	ans := maxUncrossedLines(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
