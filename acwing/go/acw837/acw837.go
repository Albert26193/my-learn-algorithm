package main

import (
	"bufio"
	"fmt"
	"os"
)

var N = 100010
var fa = make([]int, N)
var size = make([]int, N)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n, m int
	fmt.Fscan(in, &n, &m)

	initUnion(n)
	for i := 0; i < m; i++ {
		var command string
		var a, b int
		fmt.Fscan(in, &command)
		switch command {
		case "C":
			fmt.Fscan(in, &a, &b)
			merge(a, b)
		case "Q1":
			fmt.Fscan(in, &a, &b)
			connected := (find(a) == find(b))
			if connected {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		case "Q2":
			fmt.Fscan(in, &a)
			fmt.Printf("%d\n", size[find(a)])
		}
		// fmt.Println(command, a, b)
	}
}

func initUnion(count int) {
	for i := 1; i <= count; i++ {
		fa[i] = i
		size[i] = 1
	}
}

func find(a int) int {
	if fa[a] != a {
		fa[a] = find(fa[a])
	}

	return fa[a]
}

func merge(a, b int) {
	aRoot := find(a)
	bRoot := find(b)
	if aRoot != bRoot {
		fa[aRoot] = bRoot
		size[bRoot] += size[aRoot]
	}
}
