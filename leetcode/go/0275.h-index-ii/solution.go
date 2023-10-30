// Created by Albert at 2023/10/30 13:40
// leetgo: 1.3.8
// https://leetcode.cn/problems/h-index-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// check
// 0 1 2 ... mid - 1: length = mid
// citations[mid - 1] <= citationLength - mid
func hIndex(citations []int) (ans int) {
	citationsLength := len(citations)
	left, right := 0, citationsLength

	for left < right {
		mid := (left + right + 1) / 2
		if citations[mid-1] <= citationsLength-mid {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return citationsLength - right
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

	citations := Deserialize[[]int](ReadLine(reader))
	ans := hIndex(citations)

	fmt.Println("\noutput:", Serialize(ans))
}
