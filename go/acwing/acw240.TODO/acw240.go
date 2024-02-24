package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 50050

var fa = make([]int, N+1)
var eat = make(map[string]int)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n, k int
	fmt.Fscan(in, &n, &k)
	cnt := 0
	initUnion(n)
	for i := 0; i < k; i++ {
		var d, x, y int
		fmt.Fscan(in, &d, &x, &y)

		// condition 2
		if x > n || y > n {
			cnt += 1
			continue
		}

		xr := find(x)
		yr := find(y)
		eatString := strconv.Itoa(xr) + "#" + strconv.Itoa(yr)
		revString := strconv.Itoa(yr) + "#" + strconv.Itoa(xr)
		// fmt.Println(x, y)
		// fmt.Println(eatString)
		switch d {
		case 1:
			// condition 1
			if eat[eatString] == 1 || eat[revString] == 1 {
				cnt += 1
				continue
			}
			merge(x, y)
		case 2:
			// condition 3
			if xr == yr {
				cnt += 1
				continue
			}

			// condition 1
			if eat[revString] == 1 {
				cnt += 1
				continue
			}

			eat[eatString] = 1
		}
	}

	fmt.Println(cnt)
}

func initUnion(count int) {
	for i := 1; i <= count; i++ {
		fa[i] = i
	}
}

func find(a int) int {
	if fa[a] != a {
		fa[a] = find(fa[a])
	}
	return fa[a]
}

func merge(a int, b int) {
	aRoot := find(a)
	bRoot := find(b)
	if aRoot != bRoot {
		fa[aRoot] = bRoot
	}
}
