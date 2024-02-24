// Created by Albert at 2023/09/17 22:31
// leetgo: 1.3.8
// https://leetcode.cn/problems/dungeon-game/

package main

import (
	"bufio"
	"fmt"
	"math"
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

func minn(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func calculateMinimumHP(dungeon [][]int) (ans int) {
	row := len(dungeon)
	col := len(dungeon[0])

	dp := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, col+1)
		for j := 0; j < col+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	// 此处可以看成公主真正的位置,也是 dp 的出发点
	// 出发点必须是1, 因为一旦到达0就无法进行下一步行动了
	// dp[i][j] 可以看成到 "公主" 位置所需要的最小值
	dp[row][col-1], dp[row-1][col] = 1, 1
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			minCost := minn(dp[i][j+1], dp[i+1][j])
			dp[i][j] = maxx(minCost-dungeon[i][j], 1)
		}
	}

	return dp[0][0]
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
