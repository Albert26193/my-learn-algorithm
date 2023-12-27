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

	var n int
	fmt.Fscan(in, &n)
	var stones = make([]int, 2*n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &stones[i])
		stones[n+i] = stones[i]
	}

	preSum := make([]int, 2*n+1)
	for i := 1; i <= 2*n; i++ {
		preSum[i] = preSum[i-1] + stones[i-1]
	}

	// fmt.Println(stones, preSum)

	f := make([][]int, 2*n)
	d := make([][]int, 2*n)

	for i := 0; i < 2*n; i++ {
		f[i] = make([]int, 2*n)
		d[i] = make([]int, 2*n)
		for j := 0; j < 2*n; j++ {
			f[i][j] = -0x3f3f3f3f
			d[i][j] = +0x3f3f3f3f
		}
		f[i][i] = 0
		d[i][i] = 0
	}

	for len := 2; len <= n; len++ {
		for start := 0; start+len-1 < 2*n; start++ {
			end := start + len - 1

			if len == 2 {
				f[start][end] = stones[start] + stones[end]
				d[start][end] = stones[start] + stones[end]
				continue
			}

			for k := start; k+1 <= end; k++ {
				f[start][end] = maxx(f[start][k]+f[k+1][end]+preSum[end+1]-preSum[start], f[start][end])
				d[start][end] = minn(d[start][k]+d[k+1][end]+preSum[end+1]-preSum[start], d[start][end])
			}
		}
	}

	minAns := 0x3f3f3f3f
	maxAns := -0x3f3f3f3f
	for i := 0; i < n; i++ {
		minAns = minn(minAns, d[i][i+n-1])
		maxAns = maxx(maxAns, f[i][i+n-1])
	}

	fmt.Println(minAns)
	fmt.Println(maxAns)
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

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}
