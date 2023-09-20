// Created by Albert at 2023/09/20 15:16
// leetgo: 1.3.8
// https://leetcode.cn/problems/na-ying-bi/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minCount(coins []int) (ans int) {
	allCount := 0
	for _, coin := range coins {
		heapCount := (coin-1)/2 + 1
		allCount += heapCount
	}

	return allCount
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	coins := Deserialize[[]int](ReadLine(reader))
	ans := minCount(coins)
	fmt.Println("\noutput:", Serialize(ans))
}
