package main

import (
	"bufio"
	"fmt"
	"os"
)

// redo it
func main() {
	var N int
	var F int

	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &F)

	fields := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &fields[i])
	}

	left, right := 0.0, 2000.0

	for right-left > 1e-5 {
		mid := (left + right) / 2
		if check(mid, fields, F, N) {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Println(int(right * 1000))
	return
}

func check(mid float64, fields []int, F int, N int) bool {
	sum := make([]float64, N+1)
	for i := 1; i <= N; i++ {
		sum[i] = sum[i-1] + float64(fields[i-1]) - mid
	}

	minv := 0.0

	for i, j := 0, F; j <= N; i, j = i+1, j+1 {
		minv = min(minv, sum[i])
		if sum[j]-minv >= 0 {
			return true
		}
	}

	return false
}

func min(a ...float64) float64 {
	res := a[0]
	for _, num := range a {
		if num < res {
			res = num
		}
	}

	return res
}
