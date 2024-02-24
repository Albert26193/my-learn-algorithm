// Created by Albert's server at 2024/01/21 20:41
// leetgo: dev
// https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func shipWithinDays(weights []int, days int) (ans int) {
	n := len(weights)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + weights[i-1]
	}

	var check = func(cap int) bool {
		prev, cnt := 0, 1
		for i := 0; i <= n; i++ {
			if sum[i]-sum[prev] > cap {
				prev = i - 1
				cnt += 1
			}
		}

		return cnt <= days
	}

	left, right := maxx(weights...), sum[n]
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// check(15)
	return left
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
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
	weights := Deserialize[[]int](ReadLine(in))
	days := Deserialize[int](ReadLine(in))
	ans := shipWithinDays(weights, days)

	fmt.Println("\noutput:", Serialize(ans))
}
