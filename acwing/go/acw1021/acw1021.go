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

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	f := make([]int, m+1)
	f[0] = 1

	for i := 0; i < n; i++ {
		currentMoney := arr[i]
		for k := currentMoney; k <= m; k++ {
			f[k] += f[k-currentMoney]
		}
	}

	fmt.Println(f[m])
}
