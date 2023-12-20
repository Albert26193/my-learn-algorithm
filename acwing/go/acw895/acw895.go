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

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	// fmt.Println(arr)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				f[i] = maxx(f[i], f[j]+1)
			}
		}
	}

	ans := maxx(f...)
	fmt.Printf("%d", ans)
}

func maxx(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}
