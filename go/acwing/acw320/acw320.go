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

	var arr = make([]int, 2*n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		arr[n+i] = arr[i]
	}

	fmt.Println(arr)

	f := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		f[i] = make([]int, 2*n)
		f[i][i] = 0
	}

	// fmt.Println(len(f))
	// 2 3 5 10

	for len := 2; len <= n; len++ {
		for start := 0; start+len-1 < 2*n; start++ {
			end := start + len - 1

			if len == 2 {
				f[start][end] = arr[start] * arr[end] * arr[(end+1)%n]
				continue
			}

			// 2 3 5
			for k := start; k+1 <= end; k++ {
				f[start][end] = maxx(f[start][end], f[start][k]+f[k+1][end]+arr[start]*arr[(k+1)%n]*arr[(end+1)%n])
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = maxx(ans, f[i][i+n-1])
	}

	fmt.Println(ans, f[1])
}

func maxx(nums ...int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
