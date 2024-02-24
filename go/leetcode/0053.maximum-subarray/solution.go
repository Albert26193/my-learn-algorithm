// Created by Albert at 2023/11/13 12:35
// leetgo: 1.3.8
// https://leetcode.cn/problems/maximum-subarray/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// f[i] ==> end with nums[i], max sum is f[i]

func maxx(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func maxSubArray(nums []int) (ans int) {
	numsCount := len(nums)
	if numsCount == 0 {
		return 0
	}

	f := make([]int, numsCount+1)

	f[0] = nums[0]
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if f[i-1] > 0 {
			f[i] = f[i-1] + nums[i]
		} else {
			f[i] = nums[i]
		}

		ans = maxx(ans, f[i])
	}

	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	nums := Deserialize[[]int](ReadLine(reader))
	ans := maxSubArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
