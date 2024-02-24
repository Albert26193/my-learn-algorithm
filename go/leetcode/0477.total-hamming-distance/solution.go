// Created by Albert at 2023/06/26 06:36
// leetgo: 1.3.2
// https://leetcode.cn/problems/total-hamming-distance/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func totalHammingDistance(nums []int) (ans int) {
	const binaryLength int = 32

	for i := 0; i < binaryLength; i++ {
		oneCount := 0
		zeroCount := 0
		for _, num := range nums {
			if ((num >> i) & 1) == 1 {
				oneCount += 1
			} else {
				zeroCount += 1
			}
		}

		ans += (oneCount * zeroCount)
	}
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// nums := Deserialize[[]int](ReadLine(stdin))

	nums := []int{4, 14, 2}
	ans := totalHammingDistance(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
