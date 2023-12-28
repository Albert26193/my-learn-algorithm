package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAGS: dp
// f[i][j] ==> from [0, ... , i-1], volume is j, max value is f[i][j]
func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N, V int
	fmt.Fscan(in, &N)
	fmt.Fscan(in, &V)

	var v = make([]int, N)
	var w = make([]int, N)
	var s = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i])
		fmt.Fscan(in, &w[i])
		fmt.Fscan(in, &s[i])
	}

	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, V+1)
	}

	for i := 1; i <= N; i++ {
		currentVolume := v[i-1]
		for j := 0; j <= V; j++ {
			for k := 0; k <= s[i-1]; k++ {
				if j >= k*currentVolume {
					f[i][j] = maxx(f[i][j], f[i-1][j-k*currentVolume]+k*w[i-1])
				}
			}
		}
	}

	fmt.Println(v, w, s)
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
