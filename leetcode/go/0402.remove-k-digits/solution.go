// Created by Albert's server at 2024/01/22 01:20
// leetgo: dev
// https://leetcode.cn/problems/remove-k-digits/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeKdigits(num string, k int) string {
	sta := make([]rune, 0)

	for _, ch := range num {
		for len(sta) > 0 && k > 0 && ch < sta[len(sta)-1] {
			sta = sta[:len(sta)-1]
			k--
		}

		sta = append(sta, ch)
	}

	for k > 0 {
		sta = sta[:len(sta)-1]
		k--
	}

	ans := ""

	for len(sta) > 0 {
		cur := sta[len(sta)-1]
		ans = string(cur) + ans
		sta = sta[:len(sta)-1]
	}

	// remove prefix 0
	if ans == "" {
		ans = "0"
	} else {
		firstNoZeroIndex := len(ans) - 1
		for i := 0; i < len(ans); i++ {
			if ans[i] != '0' {
				firstNoZeroIndex = i
				break
			}
		}
		ans = ans[firstNoZeroIndex:]
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

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	num := Deserialize[string](ReadLine(in))
	k := Deserialize[int](ReadLine(in))
	ans := removeKdigits(num, k)

	fmt.Println("\noutput:", Serialize(ans))
}
