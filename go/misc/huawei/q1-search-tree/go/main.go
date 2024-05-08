package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	// file, err := os.Open("./testcases.txt")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }
	//
	// in := bufio.NewReader(file)
	// defer file.Close()
	in := bufio.NewReader(os.Stdin)

	var nums []int
	input, _ := in.ReadString('\n')
	strs := strings.Fields(input)
	for _, str := range strs {
		cur, _ := strconv.Atoi(str)
		nums = append(nums, cur)
	}
	var k int
	fmt.Fscan(in, &k)

	// 1. sort
	sort.Ints(nums)
	// fmt.Println(nums)

	// 2. build
	root := buildBST(nums)

	// 3. check
	fmt.Printf("%v", "S")
	check(root, k)
}

func buildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = buildBST(nums[:mid])
	root.Right = buildBST(nums[mid+1:])
	return root
}

func check(root *TreeNode, target int) {
	if root == nil {
		fmt.Printf("%v", "N")
		return
	}

	cur := root.Val
	// 如果目标值大于根节点
	if target > cur {
		if root.Right != nil {
			fmt.Printf("%s", "R")
		}
		check(root.Right, target)
	} else if target == cur {
		fmt.Printf("%s", "Y")
	} else if root.Left != nil {
		if root.Left != nil {
			fmt.Printf("%s", "L")
		}
		check(root.Left, target)
	}
}
