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
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	var nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	// 1 2 2 1 2 1
	// 2 3 3 1 2 1   // 1 3 1
	// 2 3 4 2 3 1   // 3 5 1
	// 3 4 5 3 4 2   // 1 6 1

	// diff
	diff := make([]int, n+1)
	diff[0] = nums[0]
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	for i := 0; i < m; i++ {
		var l, r, c int
		fmt.Fscan(in, &l)
		fmt.Fscan(in, &r)
		fmt.Fscan(in, &c)

		diff[l-1] += c
		diff[r-1+1] -= c
	}

	ans := make([]int, n)
	ans[0] = diff[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + diff[i]
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", ans[i])
	}
}
