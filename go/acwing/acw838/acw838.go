package main

import (
	"bufio"
	"fmt"
	"os"
)

// top: small
func sink(arr []int, x int, limit int) {
	minIndex := x

	left := 2*x + 1
	right := 2*x + 2

	if left < limit && arr[left] < arr[minIndex] {
		minIndex = left
	}

	if right < limit && arr[right] < arr[minIndex] {
		minIndex = right
	}

	if minIndex != x {
		arr[x], arr[minIndex] = arr[minIndex], arr[x]

		sink(arr, minIndex, limit)
	}
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

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	for i := (n - 1) / 2; i >= 0; i-- {
		sink(arr, i, n)
	}

	limit := n
	for i := 0; i < m; i++ {
		fmt.Println(arr[0])
		arr[0], arr[limit-1] = arr[limit-1], arr[0]
		limit -= 1

		sink(arr, 0, limit)
	}
}
