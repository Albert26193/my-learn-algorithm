// Created by Albert's server at 2024/03/07 21:10
// leetgo: dev
// https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x {
			i++
		}
	}
	return i * 2
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
	nums := Deserialize[[]int](ReadLine(in))
	ans := maxNumOfMarkedIndices(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
