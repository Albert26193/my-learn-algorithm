package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: Redo it with heap
func quickChoose(arr []int, left int, right int, k int) int {
	if left >= right {
		return left
	}

	pivot := (left + right) / 2
	store := left

	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := left; i < right; i++ {
		if arr[i] < arr[right] {
			arr[i], arr[store] = arr[store], arr[i]
			store += 1
		}
	}

	arr[store], arr[right] = arr[right], arr[store]
	if store >= k {
		return quickChoose(arr, left, store, k)
	} else {
		return quickChoose(arr, store+1, right, k)
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

	var n, k int
	fmt.Fscan(in, &n, &k)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	ans := quickChoose(arr, 0, n-1, k-1)
	fmt.Println(arr[ans])
}
