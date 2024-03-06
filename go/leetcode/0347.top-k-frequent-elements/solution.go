// Created by Albert's server at 2024/03/06 15:40
// leetgo: dev
// https://leetcode.cn/problems/top-k-frequent-elements/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type hp [][2]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(n interface{}) {
	*h = append(*h, n.([2]int))
}

func (h *hp) Pop() interface{} {
	last := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]
	return last
}

func topKFrequent(nums []int, k int) (ans []int) {
	mp := make(map[int]int)

	for _, num := range nums {
		mp[num] += 1
	}

	toSort := make([][2]int, 0)
	for k, v := range mp {
		toSort = append(toSort, [2]int{k, v})
	}

	h := &hp{}
	heap.Init(h)

	for _, comb := range toSort {
		heap.Push(h, comb)
	}

	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).([2]int)[0])
	}

	return
}

// return pivot's index

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
	ans := topKFrequent(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
