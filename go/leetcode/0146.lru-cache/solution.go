// Created by Albert's server at 2024/02/17 19:47
// leetgo: dev
// https://leetcode.cn/problems/lru-cache/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type entry struct {
	key, value int
}

type LRUCache struct {
	capacity int
	list     *list.List // 双向链表
	cached   map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cached:   map[int]*list.Element{},
	}
}

func (c *LRUCache) Get(key int) int {
	if _, has := c.cached[key]; !has {
		return -1
	}

	node := c.cached[key]
	c.list.MoveToFront(node)
	return node.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if node, has := c.cached[key]; has {
		node.Value = entry{
			key:   key,
			value: value,
		}
		c.list.MoveToFront(node)
	} else {
		c.cached[key] = c.list.PushFront(entry{key, value})

		if len(c.cached) > c.capacity {
			delete(c.cached, c.list.Remove(c.list.Back()).(entry).key)
		}
	}
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

	constructorParams := MustSplitArray(params[0])
	capacity := Deserialize[int](constructorParams[0])
	obj := Constructor(capacity)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "get":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Get(key))
			output = append(output, ans)
		case "put":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			value := Deserialize[int](methodParams[1])
			obj.Put(key, value)
			output = append(output, "null")
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
