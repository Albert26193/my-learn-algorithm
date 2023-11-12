// Created by Albert at 2023/04/19 12:39
// https://leetcode.cn/problems/corporate-flight-bookings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// redo it in 2023-11-12
func corpFlightBookings(bookings [][]int, n int) (ans []int) {

	diff := make([]int, n+1)
	ans = make([]int, n)

	for i := 0; i < len(bookings); i++ {
		first := bookings[i][0] - 1
		last := bookings[i][1] - 1
		seats := bookings[i][2]
		diff[first] += seats
		diff[last+1] -= seats
	}

	ans[0] = diff[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + diff[i]
	}

	return ans
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

	bookings := Deserialize[[][]int](ReadLine(reader))

	n := 5
	ans := corpFlightBookings(bookings, n)
	fmt.Println("output: " + Serialize(ans))
}
