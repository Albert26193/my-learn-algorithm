// Created by Albert"s server at 2024/03/03 00:32
// leetgo: dev
// https://leetcode.cn/problems/integer-to-roman/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func intToRoman(num int) string {
	roman := map[int][]string{
		1: {"I", "X", "C", "M"},
		2: {"II", "XX", "CC", "MM"},
		3: {"III", "XXX", "CCC", "MMM"},
		4: {"IV", "XL", "CD"},
		5: {"V", "L", "D"},
		6: {"VI", "LX", "DC"},
		7: {"VII", "LXX", "DCC"},
		8: {"VIII", "LXXX", "DCCC"},
		9: {"IX", "XC", "CM"},
	}

	digits := 0
	numStr := ""

	for num > 0 {
		if num%10 == 0 {
			num = num / 10
			digits++
			continue
		} else {
			numStr = roman[num%10][digits] + numStr
			num = num / 10
			digits++
		}
	}

	return numStr
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
	num := Deserialize[int](ReadLine(in))
	ans := intToRoman(num)

	fmt.Println("\noutput:", Serialize(ans))
}
