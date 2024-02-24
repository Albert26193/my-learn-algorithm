// Created by Albert's server at 2024/01/06 19:36
// leetgo: dev
// https://leetcode.cn/problems/cutting-ribbons/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxLength(ribbons []int, k int) (ans int) {

	var check = func(l int) bool {
		total := 0
		for _, r := range ribbons {
			count := r / l
			total += count
		}

		// fmt.Println(l, total)		
		return total >= k
	}

	left, right := 1, 100001

	for left < right {
		mid := (left + right + 1) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if check(right) {
		return right
	}

	return 0
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

	ribbons := Deserialize[[]int](ReadLine(in))
	k := Deserialize[int](ReadLine(in))
	ans := maxLength(ribbons, k)

	fmt.Println("\noutput:", Serialize(ans))
}
