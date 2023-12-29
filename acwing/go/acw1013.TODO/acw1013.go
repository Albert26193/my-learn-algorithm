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

	f := make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := n; k >= 1; k-- {
				f[k] = maxx(f[k], f[k-1]+profit[j][i])
			}
		}
	}

	fmt.Println(f[n])
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
