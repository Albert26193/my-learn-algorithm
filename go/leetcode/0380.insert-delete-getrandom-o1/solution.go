// Created by Albert's server at 2024/03/03 00:01
// leetgo: dev
// https://leetcode.cn/problems/insert-delete-getrandom-o1/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type RandomizedSet struct {
	recIndex map[int]int
	nums     []int
}

func Constructor() RandomizedSet {
	recIndex := make(map[int]int)
	nums := make([]int, 0)
	return RandomizedSet{
		recIndex: recIndex,
		nums:     nums,
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.recIndex[val]; ok {
		return false
	}

	n := len(r.nums)
	r.recIndex[val] = n
	r.nums = append(r.nums, val)

	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	id, ok := r.recIndex[val]
	if !ok {
		return false
	}

	// swap and remove
	// 1. swap to last
	last := len(r.nums) - 1
	lastVal := r.nums[last]
	r.nums[id] = lastVal
	r.recIndex[lastVal] = id

	// 2. remove val one(new swap one)
	r.nums = r.nums[:len(r.nums)-1]
	delete(r.recIndex, val)

	return true
}

func (r *RandomizedSet) GetRandom() (ans int) {
	if len(r.nums) == 0 {
		panic("error here")
	}
	index := rand.Intn(len(r.nums))
	return r.nums[index]
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
	ops := Deserialize[[]string](ReadLine(in))
	params := MustSplitArray(ReadLine(in))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "insert":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Insert(val))
			output = append(output, ans)
		case "remove":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Remove(val))
			output = append(output, ans)
		case "getRandom":
			ans := Serialize(obj.GetRandom())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
