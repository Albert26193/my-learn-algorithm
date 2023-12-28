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

	var n int
	fmt.Fscan(in, &n)

	// fmt.Println(n)
	books := []int{10, 20, 50, 100}

	// f[i] ==> max Volume is n, methods count
	f := make([]int, n+1)
	f[0] = 1
	for i := 0; i < 4; i++ {
		currentBook := books[i]
		for k := currentBook; k <= n; k++ {
			f[k] += f[k-currentBook]
		}
	}

	fmt.Println(f[n])
}
