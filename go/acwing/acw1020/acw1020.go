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

	var m, n, k int
	fmt.Fscan(in, &m, &n, &k)
	fmt.Println(m, n, k)

	var a = make([]int, k)
	var b = make([]int, k)
	var c = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}

	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			f[i][j] = 0x3f3f3f3f
		}
	}

	f[0][0] = 0

	for i := 0; i < k; i++ {
		currentO2 := a[i]
		currentN2 := b[i]

		for p := m; p >= 0; p-- {
			for q := n; q >= 0; q-- {
				f[p][q] = minn(f[p][q], f[maxx(p-currentO2, 0)][maxx(0, q-currentN2)]+c[i])
			}
		}
	}

	fmt.Println(f[m][n])
}

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}

	return res
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
