// Created by Albert at 2023/07/09 16:43
// leetgo: 1.3.2
// https://leetcode.cn/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func threeSum(nums []int) (ans [][]int) {
	n := len(nums)

	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right && right < n && left >= 0 {
			cur := nums[i] + nums[left] + nums[right]
			switch {
			case cur == 0:
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			case cur > 0:
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			case cur < 0:
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return ans
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
	// stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(in))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
