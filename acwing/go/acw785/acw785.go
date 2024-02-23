package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func quickSortV1(arr []int, left int, right int) {
	if right <= left {
		return
	}

	pivot := arr[left+rand.Intn(right-left+1)]
	big, small, equal := make([]int, 0), make([]int, 0), make([]int, 0)

	for _, num := range arr[left : right+1] {
		switch {
		case num > pivot:
			big = append(big, num)
		case num == pivot:
			equal = append(equal, num)
		case num < pivot:
			small = append(small, num)
		}
	}

	copy(arr[left:], append(small, append(equal, big...)...))
	// fmt.Println(arr, "ffffff")

	quickSortV1(arr, left, left+len(small)-1)
	quickSortV1(arr, left+len(small)+len(equal), right)
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

	quickSortV1(arr, 0, n-1)
	fmt.Println(arr)
}
