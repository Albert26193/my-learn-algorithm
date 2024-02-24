package main

import (
	"bufio"
	"fmt"
	"os"
)

// f[j] ==> sum of volume <= j, max worth is f[j]
// f[j] can be transfered from f[j - v[0]], f[j - v[1]], ... ==> O(n^2)

func main() {
	// read
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

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i])
		fmt.Fscan(in, &w[i])
	}

	// dp
	f := make([]int, V+1)
	for i := 0; i < N; i++ {
		currentCost := v[i]
		for k := currentCost; k <= V; k++ {
			f[k] = maxx(f[k-currentCost]+w[i], f[k])
		}
	}

	fmt.Println(f[V])
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
