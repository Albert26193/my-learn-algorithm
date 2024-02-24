// Created by Albert's server at 2024/01/30 13:33
// leetgo: dev
// https://leetcode.cn/problems/freedom-trail/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findRotateSteps(ring string, key string) (ans int) {

	return
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
	ring := Deserialize[string](ReadLine(in))
	key := Deserialize[string](ReadLine(in))
	ans := findRotateSteps(ring, key)

	fmt.Println("\noutput:", Serialize(ans))
}
