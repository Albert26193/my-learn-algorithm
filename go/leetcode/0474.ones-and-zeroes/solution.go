// Created by Albert at 2023/10/02 16:53
// leetgo: 1.3.8
// https://leetcode.cn/problems/ones-and-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// dp[i][j][k] = m: from index: 0 ... i - 1(length == i) max subset Length = m which can fit
// 0 equals j, 1 equals k
// dp[i][j][k] = dp[i - 1][j][k] ==> not choose strs[i - 1]
// dp[i][j][k] = dp[i - 1][j - current(0)][k - current(1)] + 1
func count(num byte, str string) int {
	cnt := 0
	if num != '0' && num != '1' {
		panic("error")
	}

	for i := 0; i < len(str); i++ {
		if str[i] == num {
			cnt += 1
		}
	}

	return cnt
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

func findMaxForm(strs []string, m int, n int) (ans int) {
	strsLength := len(strs)

	dp := make([][][]int, strsLength+1)
	for i := 0; i <= strsLength; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	// boundary: todo

	for i := 1; i <= strsLength; i++ {
		currentOneCounts := count('1', strs[i-1])
		currentZeroCounts := count('0', strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if currentZeroCounts <= j && currentOneCounts <= k {
					dp[i][j][k] = maxx(dp[i][j][k], dp[i-1][j-currentZeroCounts][k-currentOneCounts]+1)
				}
			}
		}
	}

	return dp[strsLength][m][n]
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

	strs := Deserialize[[]string](ReadLine(reader))
	m := Deserialize[int](ReadLine(reader))
	n := Deserialize[int](ReadLine(reader))
	ans := findMaxForm(strs, m, n)

	fmt.Println("\noutput:", Serialize(ans))
}
