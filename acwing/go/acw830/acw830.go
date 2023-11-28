package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAGS: monotonic stack

func main() {
	// read
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	// in := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(in, &N)

	var nums = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &nums[i])
	}

	// fmt.Println(N, nums)

	// Monotonic stack
	stack := make([]int, 0)
	rec := make(map[int]int)

	for i := 0; i < N; i++ {
		currentNum := nums[i]

		// pop
		for len(stack) > 0 && currentNum <= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// stack top is the value
		if len(stack) > 0 {
			rec[i] = stack[len(stack)-1]
		} else {
			// if can not find it
			rec[i] = -1
		}

		// push
		stack = append(stack, i)
	}

	ans := make([]int, 0)
	for i := 0; i < N; i++ {
		if rec[i] == -1 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, nums[rec[i]])
		}
	}

	for i := 0; i < N; i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", ans[i])
	}
}
