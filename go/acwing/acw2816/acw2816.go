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

	var a = make([]int, n)
	var b = make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	if n > m {
		fmt.Println("No")
		return
	}

	for i, j := 0, 0; i < n && j < m; j++ {
		if i < n && a[i] == b[j] {
			i++
		}

		if i == n {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
