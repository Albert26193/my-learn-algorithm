// Created by Albert's server at 2024/02/23 19:37
// leetgo: dev
// https://leetcode.cn/problems/grumpy-bookstore-owner/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	allSum := 0
	n := len(customers)
	sales := make([]int, n)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			allSum += customers[i]
		} else {
			sales[i] = customers[i]
		}
	}

	window := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			window += customers[i]
		}
	}

	curWindow := window
	for i := minutes; i < n; i++ {
		curWindow = curWindow - sales[i-minutes] + sales[i]
		window = maxx(window, curWindow)
	}

	return allSum + window
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	customers := Deserialize[[]int](ReadLine(in))
	grumpy := Deserialize[[]int](ReadLine(in))
	minutes := Deserialize[int](ReadLine(in))
	ans := maxSatisfied(customers, grumpy, minutes)

	fmt.Println("\noutput:", Serialize(ans))
}
