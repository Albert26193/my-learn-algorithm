// Created by Albert's server at 2024/01/03 23:09
// leetgo: dev
// https://leetcode.cn/problems/product-of-array-except-self/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func productExceptSelf(nums []int) (ans []int) {
	n := len(nums)

	// preProduct[i]=x, ===> length is i, product is x
	preProduct := make([]int, n+1)

	preProduct[0] = 1

	for i := 1; i <= n; i++ {
		preProduct[i] = preProduct[i-1] * nums[i-1]
	}

	rightProduct := 1
	ans = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = rightProduct * preProduct[i]
		rightProduct *= nums[i]
	}

	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	nums := Deserialize[[]int](ReadLine(in))
	ans := productExceptSelf(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
