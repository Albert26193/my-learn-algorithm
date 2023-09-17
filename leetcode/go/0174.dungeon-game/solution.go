// Created by Albert at 2023/09/17 22:31
// leetgo: 1.3.8
// https://leetcode.cn/problems/dungeon-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 1. DP; 2. memory search
func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func calculateMinimumHP(dungeon [][]int) (ans int) {
	return
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

	stdin := bufio.NewReader(reader)
	dungeon := Deserialize[[][]int](ReadLine(stdin))
	ans := calculateMinimumHP(dungeon)

	fmt.Println("\noutput:", Serialize(ans))
}
