// Created by Albert's server at 2023/12/19 15:33
// leetgo: dev
// https://leetcode.cn/problems/find-a-peak-element-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func maxx(nums []int) int {
	res := 0
	for i, v := range nums {
		if v > nums[res] {
			res = i
		}
	}

	return res
}

func findPeakGrid(mat [][]int) (ans []int) {
	n := len(mat)
	left, right := 0, n-1
	for left < right {
		mid := (left + right + 1) / 2
		y := maxx(mat[mid])
		if mat[mid][y] > mat[mid-1][y] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return []int{right, maxx(mat[right])}
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

	mat := Deserialize[[][]int](ReadLine(in))
	ans := findPeakGrid(mat)

	fmt.Println("\noutput:", Serialize(ans))
}
