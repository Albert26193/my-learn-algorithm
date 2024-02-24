// Created by Albert's server at 2023/12/23 18:07
// leetgo: dev
// https://leetcode.cn/problems/remove-stones-to-minimize-the-total/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	// "go/format"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type priorityQueue struct {
	sort.IntSlice
}

func (pq *priorityQueue) Less(i, j int) bool {
	return pq.IntSlice[i] > pq.IntSlice[j]
}

func (pq *priorityQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	arr := pq.IntSlice
	v := arr[len(arr)-1]
	pq.IntSlice = arr[:len(arr)-1]
	return v
}

func minStoneSum(piles []int, k int) (ans int) {
	pq := &priorityQueue{piles}
	heap.Init(pq)
	for i := 0; i < k; i++ {
		pile := heap.Pop(pq).(int)
		pile -= pile / 2
		heap.Push(pq, pile)
	}

	sum := 0

	for len(pq.IntSlice) > 0 {
		sum += heap.Pop(pq).(int)
	}

	return sum
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	piles := Deserialize[[]int](ReadLine(in))
	k := Deserialize[int](ReadLine(in))
	ans := minStoneSum(piles, k)

	fmt.Println("\noutput:", Serialize(ans))
}
