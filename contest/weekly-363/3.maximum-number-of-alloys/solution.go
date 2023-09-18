// Created by Albert at 2023/09/18 10:56
// leetgo: 1.3.8
// https://leetcode.cn/problems/maximum-number-of-alloys/
// https://leetcode.cn/contest/weekly-contest-363/problems/maximum-number-of-alloys/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func arraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}

func maxx(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	binarySearchAlloy := func(currentMachine int) int {
		currentComposition := composition[currentMachine]

		left := 0
		right := arraySum(stock) + budget

		// 左区间的右端点
		for left < right {
			mid := (left + right + 1) / 2
			sumCost := 0
			for i := range currentComposition {
				required := mid * currentComposition[i]

				if required <= stock[i] {
					continue
				}

				sumCost += cost[i] * (required - stock[i])
			}

			if sumCost <= budget {
				left = mid
			} else {
				right = mid - 1
			}
		}

		return right
	}

	maxAlloyCount := 0
	for i := range composition {
		maxAlloyCount = maxx(binarySearchAlloy(i), maxAlloyCount)
	}

	return maxAlloyCount
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

	n := Deserialize[int](ReadLine(reader))
	k := Deserialize[int](ReadLine(reader))
	budget := Deserialize[int](ReadLine(reader))
	composition := Deserialize[[][]int](ReadLine(reader))
	stock := Deserialize[[]int](ReadLine(reader))
	cost := Deserialize[[]int](ReadLine(reader))

	ans := maxNumberOfAlloys(n, k, budget, composition, stock, cost)

	fmt.Println("\noutput:", Serialize(ans))
}
