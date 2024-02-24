// Created by Albert at 2023/10/23 10:51
// leetgo: 1.3.8 https://leetcode.cn/problems/tuple-with-same-product/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// [a, b    1, 2]

// [a, b    1, 2]   [b, a    1, 2]   [a, b    2, 1]   [b, a    2, 1]
// ... reverse
// total: 8

func tupleSameProduct(nums []int) (ans int) {
	mp := make(map[int]int)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			mp[nums[i]*nums[j]] += 1
		}
	}

	for _, v := range mp {
		ans += v * (v - 1) * 4
	}

	return ans
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	ReadLine(reader)

	nums := Deserialize[[]int](ReadLine(reader))
	ans := tupleSameProduct(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
