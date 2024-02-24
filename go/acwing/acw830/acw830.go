package main

import (
	"bufio"
	"fmt"
	"os"
)

// TAGS: monotonic stack

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	var N int
	fmt.Fscan(in, &N)

	var nums []int
	for i := 0; i < N; i++ {
		var num int
		fmt.Fscan(in, &num)
		nums = append(nums, num)
	}

	// fmt.Println(nums)
	stack := make([]int, 0)
	rec := make(map[int]int)

	for i := 0; i < N; i++ {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rec[i] = -1
		} else {
			rec[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	for i := 0; i < N; i++ {
		out := 0

		if rec[i] == -1 {
			out = -1
		} else {
			out = nums[rec[i]]
		}

		if i > 0 {
			fmt.Printf(" %d", out)
		} else {
			fmt.Printf("%d", out)
		}
	}
}
