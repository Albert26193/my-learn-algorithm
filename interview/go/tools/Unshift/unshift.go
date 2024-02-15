package main

import "fmt"

func main() {
	nums := make([]int, 0)

	nums = append([]int{1}, nums...)
	nums = append([]int{2}, nums...)
	nums = append([]int{3}, nums...)

	fmt.Println(nums)
}

// prependNums 添加数字到切片的开头
func prependNums(nums []int, values ...int) []int {
	for _, v := range values {
		nums = append([]int{v}, nums...)
	}
	return nums
}
