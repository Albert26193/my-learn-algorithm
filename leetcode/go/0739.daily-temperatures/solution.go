// Created by Albert at 2023/10/08 12:03
// leetgo: 1.3.8
// https://leetcode.cn/problems/daily-temperatures/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func dailyTemperatures(temperatures []int) (ans []int) {
	stack := make([]int, 0)

	days := len(temperatures)
	ans = make([]int, days)

	for i := days - 1; i >= 0; i-- {
		// 1. pop
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// 2. stack Top is the element to find
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		} else {
			ans[i] = 0
		}

		// 3. push
		stack = append(stack, i)
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

	temperatures := Deserialize[[]int](ReadLine(reader))
	ans := dailyTemperatures(temperatures)

	fmt.Println("\noutput:", Serialize(ans))
}
