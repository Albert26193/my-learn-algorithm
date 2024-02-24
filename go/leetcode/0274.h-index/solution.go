// Created by Albert at 2023/10/29 13:15
// leetgo: 1.3.8
// https://leetcode.cn/problems/h-index/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// how to check
// h = mid
// 0 1 2 ... mid - 1 ==> all fit: citations[i] >= mid ==> citations[mid - 1] >= mid

func hIndex(citations []int) (ans int) {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	citationLength := len(citations)

	check := func(mid int) bool {
		if mid == 0 {
			return true
		}

		return citations[mid-1] >= mid
	}

	left, right := 0, citationLength

	for left < right {
		mid := (left + right + 1) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if check(right) {
		return right
	}

	return 0
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	reader := bufio.NewReader(file)

	defer file.Close()

	ReadLine(reader)

	citations := Deserialize[[]int](ReadLine(reader))
	ans := hIndex(citations)

	fmt.Println("\noutput:", Serialize(ans))
}
