// Created by Albert's server at 2024/02/21 02:23
// leetgo: dev
// https://leetcode.cn/problems/rectangle-area/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) (ans int) {
	lx, rx := maxx(ax1, bx1), minn(ax2, bx2)
	by, ty := maxx(ay1, by1), minn(ay2, by2)

	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	common := maxx(0, rx-lx) * maxx(0, ty-by)
	return area1 + area2 - common
}

func maxx(nums ...int) int {
	res := nums[0]

	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
}

func minn(nums ...int) int {
	res := nums[0]

	for _, v := range nums {
		if v < res {
			res = v
		}
	}

	return res
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
	ax1 := Deserialize[int](ReadLine(in))
	ay1 := Deserialize[int](ReadLine(in))
	ax2 := Deserialize[int](ReadLine(in))
	ay2 := Deserialize[int](ReadLine(in))
	bx1 := Deserialize[int](ReadLine(in))
	by1 := Deserialize[int](ReadLine(in))
	bx2 := Deserialize[int](ReadLine(in))
	by2 := Deserialize[int](ReadLine(in))
	ans := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)

	fmt.Println("\noutput:", Serialize(ans))
}
