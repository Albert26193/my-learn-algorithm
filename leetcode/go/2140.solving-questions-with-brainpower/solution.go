// Created by Albert at 2023/11/17 12:02
// leetgo: 1.3.8
// https://leetcode.cn/problems/solving-questions-with-brainpower/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// mostPoints 是解决问题的入口方法
func maxx(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func mostPoints(question [][]int) int64 {
	questionCount := len(question)
	f := make([]int, questionCount+1)
	f[0] = 0

	// forward or backward ?
	// forward
	for i, q := range question {
		f[i+1] = maxx(f[i+1], f[i])
		j := i + q[1] + 1
		if j > questionCount {
			j = questionCount
		}
		f[j] = maxx(f[j], f[i]+q[0])
	}

	return int64(f[questionCount])
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

	questions := Deserialize[[][]int](ReadLine(reader))
	ans := mostPoints(questions)

	fmt.Println("\noutput:", Serialize(ans))
}
