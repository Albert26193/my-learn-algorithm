package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 初始化一个容量为2的LRUCache
	cache := Constructor(2)

	// 插入两个元素
	cache.Put(1, 1)           // 缓存是 {1=1}
	cache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(cache.Get(1)) // 返回 1

	// 因为缓存容量是2，添加这个元素会使得键为2的元素被淘汰
	cache.Put(3, 3)           // 缓存是 {1=1, 3=3}
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)

	// 添加另一个元素，会导致键为1的元素被淘汰
	cache.Put(4, 4)           // 缓存是 {4=4, 3=3}
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回 3
	fmt.Println(cache.Get(4)) // 返回 4
}

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
	if node, has := c.cached[key]; has {
		c.list.MoveToFront(node)
		el := node.Value.(entry).value
		return el
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node, has := c.cached[key]; has {
		c.list.MoveToFront(node)
		node.Value = entry{key, value}
	} else {
		newNode := c.list.PushFront(entry{key, value})
		c.cached[key] = newNode
		if len(c.cached) > c.capacity {
			oldest := c.list.Back()
			c.list.Remove(oldest)
			delete(c.cached, oldest.Value.(entry).key)
		}
	}
}
