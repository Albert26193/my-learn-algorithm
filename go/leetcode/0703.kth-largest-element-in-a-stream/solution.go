// Created by Albert's server at 2024/03/06 20:53
// leetgo: dev
// https://leetcode.cn/problems/kth-largest-element-in-a-stream/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type KthLargest struct {

}

func Constructor(k int, nums []int) KthLargest {

	return KthLargest{}
}

func (k *KthLargest) Add(val int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	constructorParams := MustSplitArray(params[0])
	k := Deserialize[int](constructorParams[0])
	nums := Deserialize[[]int](constructorParams[1])
	numsSize := Deserialize[int](constructorParams[2])
	obj := Constructor(k, nums, numsSize)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "add":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Add(val))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
