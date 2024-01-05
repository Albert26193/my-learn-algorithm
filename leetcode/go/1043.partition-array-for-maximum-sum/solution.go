// Created by Albert's server at 2024/01/05 16:31
// leetgo: dev
// https://leetcode.cn/problems/partition-array-for-maximum-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// subArray's length can be different
// so length is within [1, k]
// TODO: redo it
func maxSumAfterPartitioning(arr []int, k int) (ans int) {
	n := len(arr)

	f := make([]int, n)
	copy(f, arr)

	// f[i]=m, end with arr[i], answer is m

	// f[i] = f[i-1] + 1 * arr[i]
	// f[i] = f[i-2] + 2 * maxx(arr[i], arr[i-1])
	// f[i] = f[i-k] + k * maxx(arr[i], arr[i-1], arr[i-k])

	for i := 0; i < n; i++ {
		curMax := -1
		for l := 1; l <= k && i-l+1 >= 0; l++ {
			curMax = maxx(curMax, arr[i-l+1])
			if i >= l {
				f[i] = maxx(f[i], f[i-l]+l*curMax)
			} else {
				f[i] = maxx(f[i], l*curMax)
			}
		}
	}

	fmt.Println(f)

	return f[n-1]
}

func maxx(nums ...int) int {
	res := nums[0]

	for _, v := range nums {
		if res < v {
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
	arr := Deserialize[[]int](ReadLine(in))
	k := Deserialize[int](ReadLine(in))
	ans := maxSumAfterPartitioning(arr, k)

	fmt.Println("\noutput:", Serialize(ans))
}
