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
	days := len(temperatures)
	stack := []int{}
	ans = make([]int, len(temperatures))

	for i := days - 1; i >= 0; i-- {
		for (len(stack) > 0) && (temperatures[i] >= temperatures[stack[len(stack)-1]]) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}

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
