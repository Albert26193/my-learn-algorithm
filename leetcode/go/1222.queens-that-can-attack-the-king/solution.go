// Created by Albert at 2023/09/14 11:47
// leetgo: 1.3.8
// https://leetcode.cn/problems/queens-that-can-attack-the-king/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func queensAttacktheKing(queens [][]int, king []int) (ans [][]int) {

	checkRange := func(x, y int) bool {
		if (x < 0) || (x >= 8) {
			return false
		}

		if (y < 0) || (y >= 8) {
			return false
		}

		return true
	}

	xKing := king[0]
	yKing := king[1]

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	queensMap := make([]int, 8*8)
	for _, queen := range queens {
		position := queen[0]*8 + queen[1]
		queensMap[position] = 1
	}

	// fmt.Println(queensMap)

	for _, direction := range directions {
		x := xKing
		y := yKing
		for checkRange(x, y) {
			if queensMap[x*8+y] == 1 {
				ans = append(ans, []int{x, y})
				break
			}
			x += direction[0]
			y += direction[1]
			// fmt.Println([]int{x, y})
		}
	}

	return
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

	queens := Deserialize[[][]int](ReadLine(reader))
	king := Deserialize[[]int](ReadLine(reader))
	ans := queensAttacktheKing(queens, king)

	fmt.Println("\noutput:", Serialize(ans))
}
