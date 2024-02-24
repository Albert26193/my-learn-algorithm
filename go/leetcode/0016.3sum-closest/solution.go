// Created by Albert at 2023/10/06 13:37
// leetgo: 1.3.8
// https://leetcode.cn/problems/3sum-closest/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func threeSumClosest(nums []int, target int) (ans int) {
	numsLength := len(nums)
	minDistance := math.MaxInt32

	sort.Slice(nums, func(a int, b int) bool {
		return nums[a] < nums[b]
	})

	for i := range nums {
		start := i + 1
		end := numsLength - 1
		tempNum := nums[i]

		for start < end {
			tempSum := nums[start] + nums[end] + tempNum
			tempDistance := tempSum - target
			tempAbsDistance := int(math.Abs(float64(tempDistance)))
			if tempDistance != 0 {
				if tempAbsDistance < minDistance {
					minDistance = tempAbsDistance
					ans = tempSum
					// fmt.Println(tempSum, start, end)
				}

				if tempDistance > 0 {
					end -= 1
				}

				if tempDistance < 0 {
					start += 1
				}
			} else {
				return target
			}
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
	target := Deserialize[int](ReadLine(reader))
	ans := threeSumClosest(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
