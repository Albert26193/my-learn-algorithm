// Created by Albert at 2023/10/07 16:37
// leetgo: 1.3.8
// https://leetcode.cn/problems/online-stock-span/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type StockSpanner struct {
	stack [][2]int
	index int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (s *StockSpanner) Next(price int) (ans int) {
	s.index++

	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}

	s.stack = append(s.stack, [2]int{s.index, price})

	return s.index - s.stack[len(s.stack)-2][0]
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

	ops := Deserialize[[]string](ReadLine(reader))
	params := MustSplitArray(ReadLine(reader))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "next":
			methodParams := MustSplitArray(params[i])
			price := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Next(price))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
