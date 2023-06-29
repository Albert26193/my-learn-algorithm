// Created by Albert at 2023/06/29 13:00
// leetgo: 1.3.2
// https://leetcode.cn/problems/reconstruct-a-2-row-binary-matrix/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// 1. check the sum of colsum and (upper + lower) is equal?
// 2. firstly fill the upper row, then lower row
func reconstructMatrix(upper int, lower int, colsum []int) (ans [][]int) {
	sumOfCols := 0
	for _, colValue := range colsum {
		if !(0 <= colValue && colValue <= 2) {
			return [][]int{}
		}

		sumOfCols += colValue
	}

	if sumOfCols != upper+lower {
		return [][]int{}
	}

	ans = make([][]int, 2)
	for i := range ans {
		ans[i] = make([]int, len(colsum))
	}

	for i, colValue := range colsum {
		if colValue == 0 {
			continue
		} else if colValue == 2 {
			ans[0][i] = 1
			upper -= 1
			ans[1][i] = 1
			lower -= 1
		}
	}

	for i, colValue := range colsum {
		if colValue == 1 {
			if upper > 0 {
				ans[0][i] = 1
				upper -= 1
			} else if lower > 0 {
				ans[1][i] = 1
				lower -= 1
			}
		}
	}

	if !(upper == 0 && lower == 0) {
		return [][]int{}
	}
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// upper := Deserialize[int](ReadLine(stdin))
	// lower := Deserialize[int](ReadLine(stdin))
	// colsum := Deserialize[[]int](ReadLine(stdin))

	upper := 5
	lower := 5
	colsum := []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}
	ans := reconstructMatrix(upper, lower, colsum)

	fmt.Println("\noutput:", Serialize(ans))
}
