package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var p []int
var ww []int
var cc []int

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)

	var n, m, w int
	fmt.Fscan(in, &n, &m, &w)

	initUnion(n)

	for i := 1; i <= n; i++ {
		var c, d int
		fmt.Fscan(in, &c, &d)
		cc[i] = c
		ww[i] = d
	}

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		union(u, v)
	}

	defer file.Close()

	f := make([]int, w+1)
	for i := 1; i <= n; i++ {
		if find(i) != i {
			continue
		}

		currentCost := cc[i]
		for j := w; j >= currentCost; j-- {
			f[j] = maxx(f[j-currentCost]+ww[i], f[j])
		}
	}

	fmt.Println(f[w])
}

func initUnion(n int) {
	p = make([]int, n+1)
	ww = make([]int, n+1)
	cc = make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}

	return p[x]
}

func union(x, y int) {
	xr := find(x)
	yr := find(y)

	if xr == yr {
		return
	}
	p[xr] = yr

	ww[yr] += ww[xr]
	cc[yr] += cc[xr]
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
}
