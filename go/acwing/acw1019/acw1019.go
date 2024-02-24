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

	var v = make([]int, n)
	var w = make([]int, n)
	var s = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i], &w[i], &s[i])
	}

	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		currentCost := v[i]
		for j := 0; j < s[i]; j++ {
			for k := m; k >= currentCost; k-- {
				f[k] = maxx(f[k], f[k-currentCost]+w[i])
			}
		}
	}

	fmt.Println(f[m])
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
