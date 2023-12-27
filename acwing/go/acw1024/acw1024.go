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

	var v, n int
	fmt.Fscan(in, &v, &n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	fmt.Println(arr)

	// f[i] = x; max Volume is i, max profit is x
	f := make([]int, v+1)
	for i := 0; i < n; i++ {
		currentCost := arr[i]
		for k := v; k >= currentCost; k-- {
			f[k] = maxx(f[k], f[k-currentCost]+arr[i])
		}
	}

	fmt.Println(v - f[v])
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
