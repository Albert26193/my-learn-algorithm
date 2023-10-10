// Created by Albert at 2023/10/10 16:55
// leetgo: 1.3.8
// https://leetcode.cn/problems/movement-of-robots/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// TODO: redo it
func getMod(num int) int {
	const mod int = 1e9 + 7
	num = (num%mod + mod) % mod
	return num
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func sumDistance(nums []int, s string, d int) (ans int) {
	for i := range nums {
		if s[i] == 'R' {
			nums[i] += d
		} else if s[i] == 'L' {
			nums[i] -= d
		}
	}

	numsLength := len(nums)
	sort.Ints(nums)
	// TLE
	/*
	   for i := range nums {
	           for j := i + 1; j < numsLength; j++ {
	               ans += getMod(abs(nums[j] - nums[i]))
	               ans = getMod(ans)
	           }
	   }
	*/

	// 0  1  2

	// 0-1 : 2 times
	// 1-2 : 1 time
	for i := 1; i < numsLength; i++ {
		ans += getMod(getMod(nums[i]-nums[i-1]) * getMod(numsLength-i) * i)
		ans = getMod(ans)
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
	s := Deserialize[string](ReadLine(reader))
	d := Deserialize[int](ReadLine(reader))
	ans := sumDistance(nums, s, d)

	fmt.Println("\noutput:", Serialize(ans))
}
