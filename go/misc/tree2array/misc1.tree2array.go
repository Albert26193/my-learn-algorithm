package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func fa(x int) int {
	return (x - 1) / 2
}

func left(x int) int {
	return 2*x + 1
}

func right(x int) int {
	return 2*x + 2
}

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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Println(x, y)
		index := int(math.Pow(2, float64(x))) + y - 1
		fmt.Println(index, fa(index), left(index), right(index))
	}
	fmt.Println(a)
}
