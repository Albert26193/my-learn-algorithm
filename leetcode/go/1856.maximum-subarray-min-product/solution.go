// Created by Albert at 2023/06/29 13:55
// leetgo: 1.3.2
// https://leetcode.cn/problems/maximum-subarray-min-product/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
// TODO: monotonic stack + prefix sum + multiplication principle

// nums: input array
// boundaryRecords[x][0] = the INDEX of first LEFT number which is less than nums[x]
// boundaryRecords[x][1] = the INDEX of first RIGHT number which is less than nums[x]
func findLeftAndRightBoundary(nums []int) [][2]int {
	numsLength := len(nums)
	boundaryRecords := make([][2]int, numsLength)

	leftMonotonicStack := make([]int, 0)
	rightMonotonicStack := make([]int, 0)

	for i := 0; i < numsLength; i++ {
		currentLeftNum := nums[i]
		currentRightNum := nums[numsLength-1-i]

		// 5 1 3
		// 5 --> 5,1 --> 5,3
		// 1 is the first left number of 3
		for len(leftMonotonicStack) > 0 && (currentLeftNum <= nums[leftMonotonicStack[len(leftMonotonicStack)-1]]) {
			leftMonotonicStack = leftMonotonicStack[0 : len(leftMonotonicStack)-1]
		}

		// when insert the map, currentLeftNum > stackTop, so stackTop is the nearest ele for currentNum
		if len(leftMonotonicStack) > 0 {
			boundaryRecords[i][0] = leftMonotonicStack[len(leftMonotonicStack)-1]
		} else {
			boundaryRecords[i][0] = -1
		}

		leftMonotonicStack = append(leftMonotonicStack, i)

		for (len(rightMonotonicStack) > 0) && (currentRightNum <= nums[rightMonotonicStack[len(rightMonotonicStack)-1]]) {
			rightMonotonicStack = rightMonotonicStack[0 : len(rightMonotonicStack)-1]
		}
		if len(rightMonotonicStack) > 0 {
			boundaryRecords[numsLength-i-1][1] = rightMonotonicStack[len(rightMonotonicStack)-1]
		} else {
			boundaryRecords[numsLength-i-1][1] = numsLength
		}

		rightMonotonicStack = append(rightMonotonicStack, numsLength-1-i)
	}
	return boundaryRecords
}

func customSum(prefixSum []int, startIndex int, endIndex int) int {
	// 0 1 2 : get 1 + 2
	// prefixSum[2 + 1] - prefixSum[1]
	return prefixSum[endIndex+1] - prefixSum[startIndex]
}

func customMax(a int, b int) int {
	if a < b {
		return b
	}

	return a
}

func maxSumMinProduct(nums []int) (ans int) {
	ans = -1
	const mod int = 1e9 + 7
	records := findLeftAndRightBoundary(nums)
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	for i, num := range nums {
		eligibleLeftIndex := records[i][0] + 1
		eligibleRightIndex := records[i][1] - 1
		currentSum := customSum(prefixSum, eligibleLeftIndex, eligibleRightIndex)
		ans = customMax(num*currentSum, ans)
	}

	ans = (ans%mod + mod) % mod
	return
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// nums := Deserialize[[]int](ReadLine(stdin))

	nums := []int{1, 2, 3, 2}
	ans := maxSumMinProduct(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
