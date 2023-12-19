package main

import (
	"bufio"
	"fmt"
	"os"
)

func partSort(left int, right int, arr []int, index int) int {
	num := arr[index]

	arr[index], arr[right] = arr[right], arr[index]

	store := left
	for i := left; i < right; i++ {
		if arr[i] < num {
			arr[i], arr[store] = arr[store], arr[i]
			store += 1
		}
	}

	arr[store], arr[right] = arr[right], arr[store]

	return store
}

func quickSort(arr []int, left int, right int) {
	if right <= left {
		return
	}

	pivotIndex := (left + right) / 2

	newPivotIndex := partSort(left, right, arr, pivotIndex)
	quickSort(arr, left, newPivotIndex-1)
	quickSort(arr, newPivotIndex+1, right)
}

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

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	quickSort(arr, 0, n - 1)
	fmt.Println(arr)
}
