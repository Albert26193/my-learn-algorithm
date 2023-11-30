package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n, m int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	uf := newUnionFind(n + 1)

	for i := 0; i < m; i++ {
		var ex string
		var a, b int
		fmt.Fscan(in, &ex)
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)
		if ex == "M" {
			uf.union(a, b)
		} else if ex == "Q" {
			if uf.connected(a, b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}

}

type unionFind struct {
	parent []int
	size   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	count := n

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &unionFind{
		parent: parent,
		size:   size,
		count:  count,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *unionFind) union(x int, y int) bool {
	xRoot := uf.find(x)
	yRoot := uf.find(y)

	if xRoot == yRoot {
		return false
	}

	if uf.size[xRoot] < uf.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}

	uf.parent[yRoot] = xRoot
	uf.size[xRoot] += uf.size[yRoot]
	uf.count -= 1

	return true
}

func (uf *unionFind) connected(x int, y int) bool {
	if uf.find(x) == uf.find(y) {
		return true
	}

	return false
}
