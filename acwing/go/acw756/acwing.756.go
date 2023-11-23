package main

import (
	"bufio"
	"fmt"
	"os"
)

// redo it
func main() {
	var n, m int

	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	ans := make([][]int, n)

	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}

	// 1 2 3
	// 8 9 4
	// 7 6 5

	left, right, top, bottom := 0, m-1, 0, n-1

	insertNum := 1

	for (left <= right) && (top <= bottom) {
		for i := left; i <= right; i++ {
			ans[top][i] = insertNum
			insertNum += 1
		}

		for i := top + 1; i <= bottom; i++ {
			ans[i][right] = insertNum
			insertNum += 1
		}

		for i := right - 1; i >= left && top < bottom; i-- {
			ans[bottom][i] = insertNum
			insertNum += 1
		}

		for i := bottom - 1; i > top && left < right; i-- {
			ans[i][left] = insertNum
			insertNum += 1
		}

		left += 1
		right -= 1
		top += 1
		bottom -= 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", ans[i][j])
		}
		fmt.Printf("\n")
	}
	return
}
