// Created by Albert's server at 2024/02/23 13:08
// leetgo: dev
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// solution1: quick select
/*
func findKthLargest(nums []int, k int) (ans int) {
	// n := len(nums)
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	pivot := nums[rand.Intn(len(nums))]
	big, small, equal := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num > pivot {
			big = append(big, num)
		} else if num < pivot {
			small = append(small, num)
		} else {
			equal = append(equal, num)
		}
	}

	// k in big
	if k <= len(big) {
		return quickSelect(big, k)
	}

	// k in small
	if len(nums)-len(small) < k {
		return quickSelect(small, k-len(nums)+len(small))
	}

	// k in large
	return pivot
}
*/

// solution2 : heap
type hp []int

// implement 5 methods for heap
func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(n any) {
	*h = append(*h, n.(int))
}

func (h *hp) Pop() any {
	num := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return num
}

func findKthLargest(nums []int, k int) (ans int) {
	h := &hp{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
	}

	for i := 0; i < k; i++ {
		ans = heap.Pop(h).(int)
	}
	// fmt.Println(h[k-1])

	return
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
	nums := Deserialize[[]int](ReadLine(in))
	k := Deserialize[int](ReadLine(in))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
