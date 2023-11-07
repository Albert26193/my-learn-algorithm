// Created by Albert at 2023/10/02 16:52
// leetgo: 1.3.8
// https://leetcode.cn/problems/coin-change-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it,
// first redo: 2023-11-07

/*
	f[i] = f[i -1]
*/

func change(amount int, coins []int) (ans int) {
	coinsCount := len(coins)
	f := make([]int, amount+1)

	f[0] = 1

	for i := 1; i <= coinsCount; i++ {
		currentCoin := coins[i-1]
		for j := 1; j <= amount; j++ {
			if j >= currentCoin {
				f[j] += f[j-currentCoin]
			}
		}
	}

	return f[amount]
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

	amount := Deserialize[int](ReadLine(reader))
	coins := Deserialize[[]int](ReadLine(reader))
	ans := change(amount, coins)

	fmt.Println("\noutput:", Serialize(ans))
}
