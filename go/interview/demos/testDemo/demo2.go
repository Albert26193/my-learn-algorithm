package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Join(a, b int) []int {
	ans := make([]int, 0)
	for i := 0; i < a; i++ {
		ans = append(ans, b)
	}

	return ans
}

func main() {
	fmt.Println("this is demo2")
}
