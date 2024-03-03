// Created by Albert's server at 2024/03/03 00:43
// leetgo: dev
// https://leetcode.cn/problems/roman-to-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func romanToInt(s string) (ans int) {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return roman[s[0:1]]
	}

	if roman[string(s[0])] >= roman[string(s[1])] {
		return romanToInt(s[1:]) + roman[s[0:1]]
	} else {
		return romanToInt(s[1:]) - roman[s[0:1]]
	}

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
	s := Deserialize[string](ReadLine(in))
	ans := romanToInt(s)

	fmt.Println("\noutput:", Serialize(ans))
}
