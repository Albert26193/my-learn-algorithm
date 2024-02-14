// Created by Albert's server at 2024/01/22 00:43
// leetgo: dev
// https://leetcode.cn/problems/maximum-swap/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSwap(num int) (ans int) {
	copyNum := num
	numList := make([]int, 0)

	for copyNum != 0 {
		cur := copyNum % 10
		numList = append([]int{cur}, numList...)
		copyNum /= 10
	}

	n := len(numList)
	maxIndex := rightMax(numList...)

	for i := 0; i < n; i++ {
		if numList[i] < numList[maxIndex[i]] {
			numList[i], numList[maxIndex[i]] = numList[maxIndex[i]], numList[i]
			break
		}
	}

	// fmt.Println(numList)
	for _, v := range numList {
		ans *= 10
		ans += v
	}

	return ans
}

// right[i] = j ==> index >= i, max index is j
func rightMax(nums ...int) []int {
	n := len(nums)
	maxIndex := n - 1
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
		res[i] = maxIndex
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
	num := Deserialize[int](ReadLine(in))
	ans := maximumSwap(num)

	fmt.Println("\noutput:", Serialize(ans))
}
