package main

import (
	"bufio"
	"fmt"
	"os"
)

// f[i][j] = p, [i, ..., j] ==> p
// f[i][j] = f[i][k] + f[k+1][j] ==> k: [i+1, .. j-1]
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

	var stones = make([]int, n)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &stones[i])
		preSum[i+1] = preSum[i] + stones[i]
	}

	// fmt.Println(stones, preSum)

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = 0x3f3f3f3f
		}
	}

	// boundary
	for i := 0; i < n; i++ {
		f[i][i] = 0
	}

	for len := 2; len <= n; len++ {
		for start := 0; start+len-1 < n; start++ {
			end := start + len - 1

			if len == 2 {
				f[start][end] = stones[start] + stones[end]
				continue
			}

			for k := start; k <= end-1; k++ {
				f[start][end] = minn(f[start][k]+f[k+1][end]+preSum[end+1]-preSum[start], f[start][end])
			}
		}
	}

	fmt.Println(f[0][n-1])
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
