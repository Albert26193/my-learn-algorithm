package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)

	for T > 0 {
		T--
		var N int
		fmt.Fscan("%d", &N)

		shops := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(in, &shops[i])
		}

		if N == 1 {
			fmt.Println(shops[0])
			continue
		} else if N > 1 {
			dp := make([]int, N+1)
			dp[0] = 0
			dp[1] = shops[0]
			for i := 2; i < N+1; i++ {
				dp[i] = max(dp[i-2]+shops[i-1], dp[i-1])
			}

			fmt.Println(dp[N])
		}

	}

	return
}
