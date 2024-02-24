// Created by Albert at 2023/11/10 15:15
// leetgo: 1.3.8
// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func successfulPairs(spells []int, potions []int, success int64) (ans []int) {
	sort.Ints(potions)
	getCount := func(target int) int {
		left, right := 0, len(potions)-1

		for left < right {
			mid := (left + right) / 2
			if int64(target*potions[mid]) >= int64(success) {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if int64(target*potions[left]) >= int64(success) {
			return len(potions) - left
		}

		return 0
	}

	spellsCount := len(spells)

	ans = make([]int, spellsCount)
	for i := range spells {
		ans[i] = getCount(spells[i])
	}
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
	spells := Deserialize[[]int](ReadLine(reader))
	potions := Deserialize[[]int](ReadLine(reader))
	success := Deserialize[int64](ReadLine(reader))
	ans := successfulPairs(spells, potions, success)

	fmt.Println("\noutput:", Serialize(ans))
}
