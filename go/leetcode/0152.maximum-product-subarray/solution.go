// Created by Albert at 2023/11/13 21:41
// leetgo: 1.3.8
// https://leetcode.cn/problems/maximum-product-subarray/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// f[i] ==> end with nums[i], max product is f[i]
func maxx(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func minn(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

func maxProduct(nums []int) (ans int) {
	numsCount := len(nums)
	f := make([]int, numsCount)
	g := make([]int, numsCount)

	f[0] = nums[0]
	g[0] = nums[0]
	ans = nums[0]

	for i := 1; i < numsCount; i++ {
		f[i] = maxx(f[i-1]*nums[i], nums[i], g[i-1]*nums[i])
		g[i] = minn(g[i-1]*nums[i], nums[i], f[i-1]*nums[i])
		ans = maxx(ans, f[i])
	}

	return
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

	nums := Deserialize[[]int](ReadLine(reader))
	ans := maxProduct(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
