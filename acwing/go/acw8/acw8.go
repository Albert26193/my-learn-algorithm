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

	var N, V, M int
	fmt.Fscan(in, &N, &V, &M)

	var v, m, w = make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i], &m[i], &w[i])
	}

	// f[i][j] = x
	// max Volume is i; max Weight is j, max Profit is x
	f := make([][]int, V+1)
	for i := 0; i <= V; i++ {
		f[i] = make([]int, M+1)
	}

	for k := 0; k < N; k++ {
		currentVolume := v[k]
		currentWeight := m[k]
		for i := V; i >= currentVolume; i-- {
			for j := M; j >= currentWeight; j-- {
				f[i][j] = maxx(f[i][j], f[i-currentVolume][j-currentWeight]+w[k])
			}
		}
	}

	fmt.Println(f[V][M])
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
