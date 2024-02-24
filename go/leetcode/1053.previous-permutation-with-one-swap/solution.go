// Created by Albert at 2023/04/05 18:04
// https://leetcode.cn/problems/previous-permutation-with-one-swap/

package main

import (
	// "bufio"
	"fmt"
	"math"

	// "os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func prevPermOpt1(arr []int) (ans []int) {
	arrLen := len(arr)
	rightMin := math.MaxInt64
	for i := arrLen - 1; i >= 0; i-- {
		rightMin = int(math.Min(float64(rightMin), float64(arr[i])))
		if rightMin < arr[i] {
			rightLegalMax := math.MinInt64
			rightLegalMaxIndex := -1
			for index, value := range arr[i+1:] {
				if value >= arr[i] {
					continue
				}

				if value > rightLegalMax {
					rightLegalMax = value
					rightLegalMaxIndex = i + index + 1
				}

			}
			arr[rightLegalMaxIndex], arr[i] = arr[i], arr[rightLegalMaxIndex]
			break
		}

	}

	return arr
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// arr := Deserialize[[]int](ReadLine(stdin))
	arr := []int{1, 9, 4, 6, 7}
	arr = []int{3, 1, 1, 3}
	ans := prevPermOpt1(arr)
	fmt.Println("output: " + Serialize(ans))
}
