package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n, p, h, m int
	fmt.Fscan(in, &n, &p, &h, &m)

	fmt.Println(n, p, h, m)

	cows := make([]int, n+1)
	diff := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cows[i] = h
		diff[i] = cows[i] - cows[i-1]
	}

	rec := make(map[string]bool)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		if a == b || a == b-1 || a == b+1 {
			continue
		}

		if a > b {
			a, b = b, a
		}

		comb := strconv.Itoa(a) + "#" + strconv.Itoa(b)
		if _, ok := rec[comb]; ok {
			continue
		}
		rec[comb] = true

		// a ... b
		// a+1 ... b-1 ---> -= 1
		diff[a+1] -= 1
		diff[b] += 1
	}

	for i := 1; i <= n; i++ {
		cows[i] = cows[i-1] + diff[i]
	}

	fmt.Println(cows[1:])
}
