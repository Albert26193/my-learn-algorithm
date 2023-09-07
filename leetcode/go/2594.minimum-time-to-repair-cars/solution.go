// Created by Albert at 2023/09/07 09:55
// leetgo: 1.3.7
// https://leetcode.cn/problems/minimum-time-to-repair-cars/

package main

import (
	"fmt"
	"math"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func check(time int64, ranks []int, cars int) bool {
	sumCars := 0
	for _, rank := range ranks {
		sumCars += int(math.Sqrt(float64((time / int64(rank)))))
	}

	if sumCars >= cars {
		return true
	}

	return false
}

func binarySearch(ranks []int, cars int) int64 {
	left, right := int64(0), int64(1e15)

	for left < right {
		mid := (left + right) / 2
		if check(mid, ranks, cars) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if check(left, ranks, cars) {
		return left
	} else {
		return -1
	}
}

func repairCars(ranks []int, cars int) (ans int64) {
	ans = binarySearch(ranks, cars)
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// ranks := Deserialize[[]int](ReadLine(stdin))
	// cars := Deserialize[int](ReadLine(stdin))

	ranks := []int{5, 1, 8}
	cars := 6
	ans := repairCars(ranks, cars)

	fmt.Println("\noutput:", Serialize(ans))
}
