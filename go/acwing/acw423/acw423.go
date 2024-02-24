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

	var t, m int
	fmt.Fscan(in, &t, &m)

	cost := make([]int, m+1)
	profit := make([]int, m+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &cost[i], &profit[i])
	}

	// fmt.Println(cost, profit)

	// f[i] = x, max Volume is i, max profit is x
	f := make([]int, t+1)
	for i := 0; i < m; i++ {
		currentTime := cost[i]
		for k := t; k >= currentTime; k-- {
			f[k] = maxx(f[k], f[k-currentTime]+profit[i])
		}
	}

	fmt.Println(f[t])
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
