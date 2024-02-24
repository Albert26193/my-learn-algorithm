// Created by Albert at 2023/09/28 20:39
// leetgo: 1.3.8
// https://leetcode.cn/problems/number-of-flowers-in-full-bloom/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 统计 nums 当中 <= mark 的 数量
// 找右端点
func binarySearch(nums []int, mark int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= mark {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[right] <= mark {
		return right + 1
	} else {
		return 0
	}
}

func fullBloomFlowers(flowers [][]int, people []int) (ans []int) {
	start, end := make([]int, len(flowers)), make([]int, len(flowers))
	for i, flower := range flowers {
		start[i] = flower[0]
		end[i] = flower[1]
	}

	sort.Slice(start, func(a int, b int) bool {
		return start[a] < start[b]
	})

	sort.Slice(end, func(a int, b int) bool {
		return end[a] < end[b]
	})

	// fmt.Println(start, end)
	for _, time := range people {
		startTimeCount := binarySearch(start, time)
		endTimeCount := binarySearch(end, time-1)
		// fmt.Println(startTimeCount, endTimeCount)
		ans = append(ans, startTimeCount-endTimeCount)
	}

	return
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// skip first line
	ReadLine(reader)

	// stdin := bufio.NewReader(os.Stdin)
	flowers := Deserialize[[][]int](ReadLine(reader))
	people := Deserialize[[]int](ReadLine(reader))
	ans := fullBloomFlowers(flowers, people)

	fmt.Println("\noutput:", Serialize(ans))
	fmt.Println(binarySearch([]int{1, 3, 4, 9}, 2))
}
