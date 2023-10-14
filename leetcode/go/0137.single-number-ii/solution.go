// Created by Albert at 2023/10/14 23:47
// leetgo: 1.3.8
// https://leetcode.cn/problems/single-number-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: redo it
// bit
func singleNumber(nums []int) int {
	ans := int32(0)

	for i := 0; i < 32; i++ {
		total := int32(0)

		for _, num := range nums {
			total += ((int32(num) >> i) & 1)
		}

		if total%3 > 0 {
			ans += (1 << i)
		}
	}

	return int(ans)
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	ReadLine(reader)

	nums := Deserialize[[]int](ReadLine(reader))
	ans := singleNumber(nums)

	fmt.Println("\noutput:", ans)
}
