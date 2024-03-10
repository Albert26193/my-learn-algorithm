// Created by Albert's server at 2024/03/06 20:48
// leetgo: dev
// https://leetcode.cn/problems/find-median-from-data-stream/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MedianFinder struct {

}

func Constructor() MedianFinder {

	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int)  {

}

func (m *MedianFinder) FindMedian() (ans float64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "addNum":
			methodParams := MustSplitArray(params[i])
			num := Deserialize[int](methodParams[0])
			obj.AddNum(num)
			output = append(output, "null")
		case "findMedian":
			ans := Serialize(obj.FindMedian())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
