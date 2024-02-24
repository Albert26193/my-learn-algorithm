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

	var n, m int
	fmt.Fscan(in, &n, &m)

	var profit = make([][]int, n)
	for i := 0; i < n; i++ {
		profit[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &profit[i][j])
		}
	}

	f := make([]int, m+1)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			for k := m; k >= j; k-- {
				profit := f[k-j] + profit[i][j-1]
				if profit > f[k] {
					cnt[i] = j
					f[k] = profit
				}
			}
		}
	}

	fmt.Println(f[m])
	for i := 0; i < n; i++ {
		fmt.Println(i + 1, cnt[i])
	}
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
