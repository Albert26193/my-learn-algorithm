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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	costBall := make([]int, k)
	costBody := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &costBall[i], &costBody[i])
	}

	// f[i][j] = x, max Ball is i, max Body is j, max Profit is x
	// f[i][j] = f[i-ballCost][j-bodyCost] + 1
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}

	for i := 0; i < k; i++ {
		currentBall := costBall[i]
		currentBody := costBody[i]
		for p := n; p >= currentBall; p-- {
			for q := m; q >= currentBody; q-- {
				f[p][q] = maxx(f[p][q], f[p-currentBall][q-currentBody]+1)
			}
		}
	}

	minBody := 0x3f3f3f3f

	for i := 1; i <= m; i++ {
		if f[n][i] == f[n][m] {
			minBody = minn(minBody, i)
		}
	}

	fmt.Println(f[n][m], m-minBody)
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if res < v {
			res = v
		}
	}

	return res
}

func minn(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if res > v {
			res = v
		}
	}

	return res
}
