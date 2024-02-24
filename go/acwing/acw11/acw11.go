package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const mod = int(1e9 + 7)
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N, V int
	fmt.Fscan(in, &N, &V)

	var v = make([]int, N)
	var w = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i], &w[i])
	}

	cnt := make([]int, V+1)
	f := make([]int, V+1)

	for i := 0; i <= V; i++ {
		cnt[i] = 1
	}

	for i := 0; i < N; i++ {
		currentVolume := v[i]
		for k := V; k >= currentVolume; k-- {
			currentProfit := f[k-currentVolume] + w[i]
			if currentProfit > f[k] {
				f[k] = currentProfit
				cnt[k] = cnt[k-currentVolume]
			} else if currentProfit == f[k] {
				cnt[k] = (cnt[k] + cnt[k-currentVolume]) % mod
			}
		}
	}

	fmt.Println((cnt[V] + mod) % mod)
}
