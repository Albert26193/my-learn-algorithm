// Created by Albert's server at 2024/03/15 22:41
// leetgo: dev
// https://leetcode.cn/problems/merge-k-sorted-lists/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeKLists(lists []*ListNode) (ans *ListNode) {
	h := &hp{}

	// init
	for _, head := range lists {
		if head == nil {
			continue
		}
		*h = append(*h, head)
	}
	heap.Init(h)

	// push
	dummy := &ListNode{}
	cur := dummy

	for len(*h) > 0 {
		node := heap.Pop(h).(*ListNode)

		if node.Next != nil {
			heap.Push(h, node.Next)
		}

		cur.Next = node
		cur = cur.Next
	}

	return dummy.Next
}

type hp []*ListNode

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Len() int {
	return len(h)
}

func (h *hp) Swap(i, j int) {
	a := *h
	a[i], a[j] = a[j], a[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
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
	lists := Deserialize[[]*ListNode](ReadLine(in))
	ans := mergeKLists(lists)

	fmt.Println("\noutput:", Serialize(ans))
}
