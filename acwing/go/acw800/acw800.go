package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var n, m, x int
	fmt.Fscan(in, &n, &m, &x)

	array1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &array1[i])
	}

	//reverse
	sort.Slice(array1, func(i, j int) bool {
		return array1[j] < array1[i]
	})

	fmt.Println(array1)
	array2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &array2[i])
	}

	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && array1[i]+array2[j] < x {
			j++
		}

		if array1[i]+array2[j] == x {
			fmt.Printf("%d %d", n-1-i, j)
			return
		}
	}
}
