/*
 * @lc app=leetcode.cn id=2367 lang=golang
 *
 * [2367] 算术三元组的数目
 */
package main

import "fmt"

// @lc code=start
func arithmeticTriplets(nums []int, diff int) (ans int)  {
	vis := [301]bool{}
	for _, x := range nums {
		vis[x] = true
	}
	for _, x := range nums {
		if vis[x+diff] && vis[x+diff+diff] {
			ans++
		}
	}
	return
}

// @lc code=end

func main() {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	ans := arithmeticTriplets(nums, diff)
	fmt.Println(ans)
}
