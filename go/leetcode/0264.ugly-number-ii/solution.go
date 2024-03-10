// Created by Albert's server at 2024/03/06 20:30
// leetgo: dev
// https://leetcode.cn/problems/ugly-number-ii/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type hp []int

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Len() int {
	return len(h)
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(n interface{}) {
	*h = append(*h, n.(int))
}

func (h *hp) Pop() interface{} {
	num := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return num
}

func nthUglyNumber(n int) (ans int) {
	h := &hp{}
	heap.Init(h)

	vis := make(map[int]bool)
	factors := []int{2, 3, 5}

	heap.Push(h, 1)
	for i := 0; i <= n; i++ {
		x := heap.Pop(h).(int)

		if i == n-1 {
			return x
		}

		for _, f := range factors {
			next := x * f
			if !vis[next] {
				heap.Push(h, next)
				vis[next] = true
			}
		}
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

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	n := Deserialize[int](ReadLine(in))
	ans := nthUglyNumber(n)

	fmt.Println("\noutput:", Serialize(ans))
}
