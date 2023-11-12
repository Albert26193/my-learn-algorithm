// Created by Albert at 2023/11/12 22:08
// leetgo: 1.3.8
// https://leetcode.cn/problems/car-pooling/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1005)
	record := make([]int, 1005)

	for i := 0; i < len(trips); i++ {
		addPassengers := trips[i][0]
		from := trips[i][1]
		to := trips[i][2]

		diff[from] += addPassengers
		diff[to] -= addPassengers
	}

	for i := 0; i < 1005; i++ {
		if i == 0 {
			record[0] = diff[0]
		} else {
			record[i] = diff[i] + record[i-1]
		}

		if record[i] > capacity {
			return false
		}
	}

	// fmt.Println(record)

	return true
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

	trips := Deserialize[[][]int](ReadLine(reader))
	capacity := Deserialize[int](ReadLine(reader))
	ans := carPooling(trips, capacity)

	fmt.Println("\noutput:", Serialize(ans))
}
