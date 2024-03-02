// Created by Albert's server at 2024/02/24 14:39
// leetgo: dev
// https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func closestNodes(root *TreeNode, queries []int) (ans [][]int) {
	nums := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	for _, q := range queries {
		rightNum := findRight(nums, q)
		leftNum := findLeft(nums, q)
		ans = append(ans, []int{rightNum, leftNum})
	}

	// fmt.Println(nums)
	return
}

func findLeft(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] >= target {
		return nums[left]
	}

	return -1
}

func findRight(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[right] <= target {
		return nums[right]
	}

	return -1
}

// @lc code=end

func main() {
	file, err := os.Open("./testcases.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	in := bufio.NewReader(file)
	defer file.Close()

	ReadLine(in)
	root := Deserialize[*TreeNode](ReadLine(in))
	queries := Deserialize[[]int](ReadLine(in))
	ans := closestNodes(root, queries)

	fmt.Println("\noutput:", Serialize(ans))
}
