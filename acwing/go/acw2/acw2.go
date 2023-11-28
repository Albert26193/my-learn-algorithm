package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAGS: dp

// 01 package
// f[i][j] ==> from [0, ... i-1] choose items, volume is j, max worth is f[i][j]

// choose or not choose
// f[i][j] = f[i-1][j]
// f[i][j] = f[i-1][j - v[i-1]] + w[i-1]

func main() {
	// input
	// file, err := os.Open("./testcases.txt")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }
	//
	// in := bufio.NewReader(file)
	// defer file.Close()

	in := bufio.NewReader(os.Stdin)
	var N, V int
	fmt.Fscan(in, &N)
	fmt.Fscan(in, &V)

	var v = make([]int, N)
	var w = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i])
		fmt.Fscan(in, &w[i])
	}

	// fmt.Println(N, V, v, w)

	// core
	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, V+1)
	}

	f[0][0] = 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= V; j++ {
			f[i][j] = f[i-1][j]
			currentVolume := v[i-1]
			if j >= currentVolume {
				f[i][j] = maxx(f[i][j], f[i-1][j-currentVolume]+w[i-1])
			}
		}
	}

	fmt.Println(f[N][V])
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
