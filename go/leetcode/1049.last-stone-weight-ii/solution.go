// Created by Albert's server at 2024/01/06 21:54
// leetgo: dev
// https://leetcode.cn/problems/last-stone-weight-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func getSum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func lastStoneWeightII(stones []int) (ans int) {
	// add -,
	// part stones into two parts
	// the lower part --> try best to get close to sum/2

	summ := getSum(stones...)
	target := summ / 2
	fmt.Println(summ, target)

	f := make([]bool, target+1)

	n := len(stones)
	f[0] = true
	for i := 0; i < n; i++ {
		currentVolume := stones[i]
		for j := target; j >= currentVolume; j-- {
			f[j] = (f[j-currentVolume] || f[j])
		}
	}

	for i := target; i >= 0; i-- {
		if f[i] == true {
			return (summ - i - i)
		}
	}

	return -1
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
	stones := Deserialize[[]int](ReadLine(in))
	ans := lastStoneWeightII(stones)

	fmt.Println("\noutput:", Serialize(ans))
}
