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

	fmt.Println(n, m)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	f := make([]int, n+1)
	f[0] = 1
	for i := 0; i < n; i++ {
		currentNum := arr[i]
		for k := m; k >= currentNum; k-- {
			f[k] += f[k-currentNum]
		}
	}

	fmt.Println(f[m])
}
