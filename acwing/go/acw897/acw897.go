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
	var aa, bb string

	fmt.Fscan(in, &n, &m, &aa, &bb)
	fmt.Println(aa, bb)

	// f[i][j] = m, aa: [0,...i-1], bb: [0, ..., j-1], max common = m
	// f[i][j] = f[i-1][j] or f[i][j-1]
	// f[i][j] = f[i-1][j-1] + 1
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}

	f[0][0] = 0
	f[1][0] = 0
	f[0][1] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if aa[i-1] == bb[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = maxx(f[i-1][j], f[i][j-1])
			}
		}
	}

	fmt.Printf("%d", f[n][m])
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if res < v {
			res = v
		}
	}
	return res
}
