// Created by Albert at 2023/04/19 12:39
// https://leetcode.cn/problems/corporate-flight-bookings/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func corpFlightBookings(bookings [][]int, n int) (ans []int) {
	diff := make([]int, n+1)
	ans = make([]int, n)
	for _, singleBooking := range bookings {
		startIndex := singleBooking[0] - 1
		endIndex := singleBooking[1] - 1
		ticketCount := singleBooking[2]

		diff[startIndex] += ticketCount
		diff[endIndex+1] -= ticketCount
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			ans[0] = diff[0]
		} else {
			ans[i] = diff[i] + ans[i-1]
		}
	}

	return
}

// @lc code=end
func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// bookings := Deserialize[[][]int](ReadLine(stdin))
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5
	ans := corpFlightBookings(bookings, n)
	fmt.Println("output: " + Serialize(ans))
}
