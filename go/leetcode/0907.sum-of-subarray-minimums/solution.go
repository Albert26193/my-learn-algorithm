// Created by Albert at 2023/06/29 12:53
// leetgo: 1.3.2
// https://leetcode.cn/problems/sum-of-subarray-minimums/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7

	arrLength := len(arr)
	stack := []int{}

	recordLeft := make([]int, arrLength)
	recordRight := make([]int, arrLength)

	for i := 0; i < arrLength; i++ {
		for (len(stack) > 0) && (arr[i] < arr[stack[len(stack)-1]]) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			recordLeft[i] = -1
		} else {
			recordLeft[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := arrLength - 1; i >= 0; i-- {
		// why <= ? ==> to filter first left Num which equals arr[i]
		for (len(stack) > 0) && (arr[i] <= arr[stack[len(stack)-1]]) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			recordRight[i] = arrLength
		} else {
			recordRight[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// fmt.Println(recordLeft, recordRight)
	for i := range arr {
		sum := (arr[i]*(i-recordLeft[i])*(recordRight[i]-i)%mod + mod) % mod
		ans = ((ans+sum)%mod + mod) % mod
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

	arr := Deserialize[[]int](ReadLine(reader))
	ans := sumSubarrayMins(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
