// Created by Albert's server at 2024/02/17 19:47
// leetgo: dev
// https://leetcode.cn/problems/lru-cache/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

func initDlinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkNode{},
		head:     initDlinkNode(0, 0),
		tail:     initDlinkNode(0, 0),
		capacity: capacity,
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) (ans int) {
	if _, ok := l.cache[key]; !ok {
		return -1
	}

	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; ok {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	} else {
		node := initDlinkNode(key, value)
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	}
}

func (l *LRUCache) moveToHead(node *DLinkNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *DLinkNode) {
	node.prev = l.head
	node.next = l.head.next

	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeTail() *DLinkNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
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
