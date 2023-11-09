// Created by Albert at 2023/11/09 11:40
// leetgo: 1.3.8
// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// for every element, it will make positive contribution to total window
func countSubarrays(nums []int, k int64) (ans int64) {
	lenNums := len(nums)
	preSum := make([]int, lenNums+1)

	for i := 1; i <= lenNums; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	// [... m ... n...] => get sum of [m,n] => preSum[n+1] - preSum[m]
	// n: 2, m: 0
	fast := 0
	slow := 0
	for fast = 0; fast < lenNums; fast += 1 {
		currentSum := int64(preSum[fast+1] - preSum[slow])
		currentProduct := int64(currentSum * int64(fast-slow+1))
		for currentProduct >= k && slow <= fast {
			slow += 1
			currentSum = int64(preSum[fast+1] - preSum[slow])
			currentProduct = int64(currentSum * int64(fast-slow+1))
		}

		// fix fast
		// [slow+0, ... , fast] <= k
		// [slow+1 ... , fast] <= k
		// [slow+2 ... , fast] <= k ==> count: fast - slow + 1
		if currentProduct < k {
			ans += int64(fast - slow + 1)
		}
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
	k := Deserialize[int64](ReadLine(reader))
	ans := countSubarrays(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
